import { compressURL, decompressURL } from './url-compressor.js';

const get = id => document.getElementById(id);
const urlInput = get('url-input');
const encodedOutput = get('encoded-output');
const encodedInput = get('encoded-input');
const decodedOutput = get('decoded-output');
const status = get('status');
const compressedURLPrefix = 'HTTPS://C.SHOGO82148.COM/';
const byteLength = value => new TextEncoder().encode(value).length;

function compress() {
  resetQR();
  status.textContent = '';
  const input = urlInput.value;
  encodedOutput.value = input ? compressedURLPrefix + compressURL(input) : '';
  get('copy-encoded').disabled = !input;
  get('transfer').disabled = !input;
  get('generate-qr').disabled = !input;
  const before = byteLength(input);
  const after = byteLength(encodedOutput.value);
  get('original-size').textContent = before ? before.toLocaleString() : '—';
  get('compressed-size').textContent = before ? after.toLocaleString() : '—';
  get('change-label').textContent = after > before ? '増加率' : '削減率';
  get('change').textContent = before ? `${(Math.abs(before - after) / before * 100).toFixed(1)}%` : '—';
}

function decompress() {
  status.textContent = '';
  decodedOutput.value = '';
  get('decode-error').textContent = '';
  encodedInput.removeAttribute('aria-invalid');
  try {
    if (encodedInput.value) {
      const input = encodedInput.value;
      const payload = input.slice(0, compressedURLPrefix.length).toUpperCase() === compressedURLPrefix
        ? input.slice(compressedURLPrefix.length)
        : input;
      decodedOutput.value = decompressURL(payload);
    }
  } catch (error) {
    encodedInput.setAttribute('aria-invalid', 'true');
    get('decode-error').textContent = `復号できません。圧縮文字列を確認してください（${error.message}）。`;
  }
  get('copy-decoded').disabled = !decodedOutput.value;
}

async function copy(output) {
  try {
    await navigator.clipboard.writeText(output.value);
    status.textContent = 'コピーしました。';
  } catch {
    output.focus();
    output.select();
    status.textContent = '自動コピーできませんでした。選択されたテキストをコピーしてください。';
  }
}

// Replace image elements to keep late load/error events from updating newer results.
let qrGeneration = 0;
function resetQR() {
  qrGeneration++;
  get('qr-results').hidden = true;
  for (const id of ['original-qr', 'compressed-qr']) {
    get(id).replaceChildren();
    get(`${id}-status`).textContent = '';
  }
}

function showQR(id, data, alt) {
  const generation = qrGeneration;
  const image = new Image();
  image.alt = alt;
  image.width = 280;
  image.height = 280;
  image.referrerPolicy = 'no-referrer';
  const message = get(`${id}-status`);
  message.textContent = '生成中…';
  image.hidden = true;
  image.onload = () => {
    if (generation !== qrGeneration) return;
    image.hidden = false;
    message.textContent = '';
  };
  image.onerror = () => {
    if (generation !== qrGeneration) return;
    message.textContent = '生成できませんでした。入力の長さや接続を確認し、もう一度お試しください。';
  };
  const url = new URL('https://qr.shogo82148.com/qr');
  url.search = new URLSearchParams({ data, size: '280', format: 'svg', level: 'M' });
  get(id).replaceChildren(image);
  image.src = url.href;
}

get('generate-qr').addEventListener('click', () => {
  if (!urlInput.value) return;
  resetQR();
  get('qr-results').hidden = false;
  showQR('original-qr', urlInput.value, '元のURLのQRコード');
  showQR('compressed-qr', encodedOutput.value, '圧縮URLのQRコード');
});

urlInput.addEventListener('input', compress);
encodedInput.addEventListener('input', decompress);
get('sample').addEventListener('click', () => {
  urlInput.value = 'https://example.com/articles/url-compressor?source=playground#overview';
  compress();
});
get('clear').addEventListener('click', () => {
  urlInput.value = '';
  compress();
  urlInput.focus();
});
get('transfer').addEventListener('click', () => {
  encodedInput.value = encodedOutput.value;
  decompress();
  encodedInput.focus();
});
get('copy-encoded').addEventListener('click', () => copy(encodedOutput));
get('copy-decoded').addEventListener('click', () => copy(decodedOutput));
compress();
decompress();
