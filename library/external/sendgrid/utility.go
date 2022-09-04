// ユーティリティパッケージ
package sendgrid

import (
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// SetConv
// 置換文字列のセット
// param p : 送信対象情報の設定先
func (info *Info) SetConv(p *mail.Personalization) {
	// 送信対象
	t := info.Target

	// メール本文の置換
	p.SetSubstitution("{%name%}", t.Name)
	p.SetSubstitution("{%email%}", t.Email)
	p.SetSubstitution("{%url%}", "")
	p.SetSubstitution("{%reset_pass_url%}", "")
	p.SetSubstitution("{%unsubscribe_url%}", "")
}