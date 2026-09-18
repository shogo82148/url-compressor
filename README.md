# url-compressor

URLを固定Huffman符号とBase45で圧縮します。

## JavaScript

`url-compressor.js` は外部依存のないESモジュールです。Node.jsとブラウザの
`<script type="module">` から同じAPIを利用できます。

```js
import { compressURL, decompressURL, encode, decode } from './url-compressor.js';

const compressed = compressURL('https://example.com/path?q=hello');
console.log(decompressURL(compressed)); // https://example.com/path?q=hello

// 既存Goサーバーで利用する圧縮URL
const link = `HTTPS://COMPRESSOR.EXAMPLE/${compressed}`;

// Goのencode/decodeに対応する低水準API
const encoded = encode('example.com'); // UTF-8文字列またはUint8Arrayを受け取る
const bytes = decode(encoded);        // Uint8Arrayを返す
```

- `compressURL(url)` は先頭の `http://`、`https://` をGo版と同じ順で除去し、
  バージョン `0` で始まる圧縮文字列を返します。
- `decompressURL(data)` は圧縮文字列を受け取り、`https://` を付けて復元します。
  元のHTTPスキームは保存しません。ホスト名やパーセントエンコードの正規化は行いません。
- Huffmanテーブル、ビット順、末尾の1によるパディング、Base45、エスケープは
  Go版と互換です。復号もGo版と同様に末尾の未完の符号を無視します。
- 空のデータは `0` になります。空の圧縮文字列、未知のバージョン、不正なBase45は
  `Error` を送出します。バイナリデータの復元には `decode` を使ってください。

## テスト

Node.js 20以降とGoで実行します。npmパッケージのインストールは不要です。

```sh
npm test
go test ./...
```

`testdata/compatibility.json` は既存Go実装から生成した共通テストデータです。
全256バイト、UTF-8のURL、空入力、固定シードのバイナリ入力について、
両実装の符号化結果と復号結果を検証します。

## HTMLプレイグラウンド

PlaygroundのHTMLとJavaScriptは `go:embed` でGoバイナリーに組み込まれます。
ビルドして起動します。

```sh
go build -o url-compressor .
./url-compressor
```

ブラウザで <http://localhost:8080/> を開いてください。
実行時にHTML・JavaScriptファイルを配置する必要はありません。
`/` は従来のプレビューに代わりPlaygroundを表示し、`/<圧縮文字列>` の復号ページは引き続き利用できます。

URLの自動圧縮、バイト数と削減率の表示、圧縮文字列の復号、結果のコピーを試せます。
圧縮・復号はブラウザ内だけで処理されます。

「QRコードを表示」で、元のURLと圧縮URLのQRコードを並べて表示できます。
生成対象は [QR Studio](https://qr.shogo82148.com/) の `/qr` API に送信されます。
両方とも280px・SVG・誤り訂正レベルMです。プレイグラウンドの圧縮結果・コピー・QRコードには
`HTTPS://C.SHOGO82148.COM/` を付けたURLを使用し、サイズ比較にもプレフィックスを含めます。
復号欄はこのURL全体と、プレフィックスのない圧縮文字列の両方に対応します。
元のURLを変更すると古いQRコードは消え、再度ボタンを押すと生成されます。
