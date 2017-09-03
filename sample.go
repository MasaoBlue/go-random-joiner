package main

import "fmt"
import "os"
import "bufio"
import "time"
import "math/rand"

func main(){
  // テキスト配列を定義
  // 「:=」で変数宣言と代入を同時に行う
  words := []string{}

  // ファイルを開く
  fp, err := os.Open("words.txt")
  if err != nil {
    panic(err)
  }
  scanner := bufio.NewScanner(fp) // 1行ずつ読み取るためbufio(バッファ付きIO)を使用
  for scanner.Scan() {
    // 1行ずつ読み取ってwords配列に格納
    words = append(words, scanner.Text())
  }
  fp.Close()

  fmt.Printf("こんにちは。\n>") // 入力前のプロンプトを表示

  // 複数回実行を考慮し、ファイル名に現在時刻を入れる
  const format = "20060102_1504"
  now := time.Now()
  outputFile := "result_" + now.Format(format) + ".txt"

  // 書き込み用ファイルを開く
  fpw, err := os.Create(outputFile)
  if err != nil {
    panic(err)
  }
  w := bufio.NewWriter(fpw)

  for {
    // 入力受け付け
    var stdin string
    fmt.Scan(&stdin)

    fmt.Fprint(w, stdin + "\n")
    w.Flush()

    fmt.Printf(choice(words) + "\n>")
  }

  fpw.Close()
}

// 配列からランダム抽出
func choice(s []string) string {
  rand.Seed(time.Now().UnixNano())
  i := rand.Intn(len(s))
  return s[i]
}
