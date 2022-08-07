package main

func main() {
	var int val1
	var string val2
	err = rows.Scan(&val1, &val2) // テーブルのval1, val2というカラムの中身が取れる

}

type MyType string //エイリアスを作る

// インターフェースを定義
func (mType *myType) Scan(src interface{}) error {
  switch x := src.(type) {
    //型ごとに分岐して 実体の*mType に MyType({変換した文字列の値}) を代入。
    // 最後はnilを返す。途中エラーが起こったら戻り値はそのエラーが入る。
  }
}

// DBアクセスの結果がrowsに入っているとする
var age MyType
var err
err = rows.Scan(&id, &name, &age)
// 3番目のフィールド内容が変数ageの指すアドレス先に入る過程で、上の処理が走る

var name sql.NullString // この方法ではポインタの*不要
rows.Scan(&name)
fmt.Println(name.String) // 結果がNULLだったら""にしてくれる

type MyNullString struct {
	s sql.NullString
  }
  
  func (mns *MyNullString) Scan(value interface{}) error {
	// mns.s.Scan(value) を呼ぶだけ
  }
  
  func (mns MyNullString) String() string {
	// 文字列化するメソッド
  }
  
  func (mns MyNullString) MarshalJSON() ([]byte, error) {
	// 引数チェックの後にjson.Marshal(mns.s.String)を呼ぶ
  }
  
  func (mns *MyNullString) UnmarshalJSON(data []byte) error {
	// json.Unmarshal を呼んだ結果をmns.s.String と mns.s.Validに入れてnilを返す
  }
  
  
  // 実際に使う側
  type SampleUser struct {
	ID int64
	Name MyNullString
  }
  
  var user SampleUser
  err = row.Scan(&user.ID, &user.Name)
  // この変数をJSON化処理や標準出力に渡してもうまく動作する


  func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
	  return c.String(http.StatusOK, "Hello")
	})
	e.Logger.Fatal(e.Start(":8080"))
  }