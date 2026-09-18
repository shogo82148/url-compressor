import test from 'node:test';
import assert from 'node:assert/strict';
import { readFileSync } from 'node:fs';
import { encode, decode, compressURL, decompressURL } from './url-compressor.js';

const vectors = JSON.parse(readFileSync(new URL('./testdata/compatibility.json', import.meta.url)));

test('matches Go encoding and decoding for UTF-8, all bytes, and seeded binary inputs', () => {
  for (const { hex, encoded } of vectors) {
    const bytes = Uint8Array.from(Buffer.from(hex, 'hex'));
    assert.equal(encode(bytes), encoded, `encode ${hex}`);
    assert.deepEqual(decode(encoded), bytes, `decode ${hex}`);
  }
});

test('URL helpers match scheme removal and HTTPS restoration', () => {
  for (const url of ['example.com', '例え.jp/日本語?q=🍣', '\ufeffexample.com']) {
    assert.equal(encode(url), encode(new TextEncoder().encode(url)));
    for (const prefix of ['', 'http://', 'https://', 'http://https://']) {
      assert.equal(compressURL(prefix + url), encode(url));
      assert.equal(decompressURL(compressURL(prefix + url)), 'https://' + url);
    }
  }
  assert.equal(compressURL('HTTP://example.com'), encode('HTTP://example.com'));
  assert.equal(encode(''), '0');
  assert.deepEqual(decode('0'), new Uint8Array());
});

test('decodes all legacy escape aliases without recursively replacing', () => {
  for (const [index, char] of [...' $%*+-./:'].entries()) {
    const escaped = '.' + String.fromCharCode(65 + index);
    assert.deepEqual(decode('0' + escaped + '0'), decode('0' + char + '0'));
  }
  assert.throws(() => decode('0.GA'), /invalid value/);
});

test('rejects invalid version, Base45 length, characters, and overflow', () => {
  for (const [input, error] of [
    ['', /empty data/], ['1ABC', /unknown version/], ['00', /invalid length/],
    ['0aa', /invalid character/], ['0.J', /invalid value/],
    ['0::', /invalid value/], ['0:::', /invalid value/],
  ]) assert.throws(() => decode(input), error);
  assert.throws(() => encode([1, 2]), TypeError);
  assert.throws(() => decode(null), TypeError);
  assert.throws(() => compressURL(null), TypeError);
});

test('handles large input without overflowing the JS bitwise accumulator', () => {
  const input = 'example.com/a?x=日本語&y=%20#z'.repeat(10000);
  assert.equal(new TextDecoder().decode(decode(encode(input))), input);
});
