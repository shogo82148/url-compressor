import { encodeTable } from './huffman.js';

const alphabet = '0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ $%*+-./:';
const escapes = { ' ': '.A', '%': '.C', '*': '.D', '+': '.E', '.': '.G', '/': '.H' };
const unescapes = { A: ' ', B: '$', C: '%', D: '*', E: '+', F: '-', G: '.', H: '/', I: ':' };
const decodeTree = {};
for (const [value, [code, length]] of encodeTable.entries()) {
  let node = decodeTree;
  for (let i = length - 1; i >= 0; i--) {
    node = node[(code >>> i) & 1] ??= {};
  }
  node.value = value;
}

/** Encode a UTF-8 string or Uint8Array using the Go version-0 format. */
export function encode(data) {
  if (typeof data === 'string') data = new TextEncoder().encode(data);
  if (!(data instanceof Uint8Array)) throw new TypeError('expected a string or Uint8Array');
  const bytes = [];
  let byte = 0;
  let used = 0;
  for (const value of data) {
    const [code, length] = encodeTable[value];
    for (let i = length - 1; i >= 0; i--) {
      byte = (byte << 1) | ((code >>> i) & 1);
      if (++used === 8) {
        bytes.push(byte);
        byte = 0;
        used = 0;
      }
    }
  }
  // Pad with a prefix of the all-ones terminator to the next byte boundary.
  if (used) bytes.push((byte << (8 - used)) | ((1 << (8 - used)) - 1));
  let result = '0';
  for (let i = 0; i < bytes.length; i += 2) {
    const paired = i + 1 < bytes.length;
    let value = paired ? bytes[i] * 256 + bytes[i + 1] : bytes[i];
    result += alphabet[value % 45];
    value = Math.floor(value / 45);
    result += alphabet[value % 45];
    if (paired) result += alphabet[Math.floor(value / 45)];
  }
  return result.replace(/[ %*+./]/g, char => escapes[char]);
}

/** Decode to bytes, preserving arbitrary binary input just like Go's decode. */
export function decode(data) {
  if (typeof data !== 'string') throw new TypeError('expected a string');
  if (!data.length) throw new Error('empty data');
  if (data[0] !== '0') throw new Error(`unknown version: ${data[0]}`);
  const source = data.slice(1).replace(/\.([A-I])/g, (_, key) => unescapes[key]);
  if (source.length % 3 === 1) throw new Error('base45: invalid length');
  const bytes = [];
  for (let i = 0; i < source.length; i += 3) {
    const count = Math.min(3, source.length - i);
    let value = 0;
    for (let j = 0; j < count; j++) {
      const digit = alphabet.indexOf(source[i + j]);
      if (digit < 0) throw new Error('base45: invalid character');
      value += digit * 45 ** j;
    }
    if (value >= (count === 3 ? 65536 : 256)) throw new Error('base45: invalid value');
    if (count === 3) bytes.push(value >>> 8);
    bytes.push(value & 255);
  }
  const result = [];
  let node = decodeTree;
  for (const byte of bytes) {
    for (let i = 7; i >= 0; i--) {
      node = node[(byte >>> i) & 1];
      if (node.value !== undefined) {
        // Go converts even the full terminator (256) to a byte.
        result.push(node.value & 255);
        node = decodeTree;
      }
    }
  }
  // Like Go, ignore any incomplete final code, including padding.
  return Uint8Array.from(result);
}

/** Remove the scheme using the same two prefix removals as serveIndex. */
export function compressURL(url) {
  if (typeof url !== 'string') throw new TypeError('expected a string');
  return encode(url.replace(/^http:\/\//, '').replace(/^https:\/\//, ''));
}

/** Restore an HTTPS URL, matching the Go redirect handler. */
export function decompressURL(data) {
  return 'https://' + new TextDecoder('utf-8', { ignoreBOM: true }).decode(decode(data));
}
