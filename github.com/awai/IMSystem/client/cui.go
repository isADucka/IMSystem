/*
 * @Author: cyy 2867025942@qq.com
 * @Description: 交互界面
 */
package client

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/awai/IMSystem/client/sdk"
	"github.com/gookit/color"
	"github.com/rocket049/gocui"
)

func init(){
	rand.Seed(time.Now().UnixNano())

}

var (
	buf string   // 写入文件的数据
	chat *sdk.Chat   //Chat对象
	pos int  //与事件有关，用来辅助查找输入框的上下条数据
)

/**
* cmd/client.go调用的
 */
func RunMain() {
	//①创建chat的核心对象,先写死数据
	chat = sdk.NewChat("127.0.0.1:8080", "awai", "123123123", "123123")
	//② 创建gui界面
	cui, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	//光标,鼠标和编码
	cui.Cursor = true
	cui.Mouse = false
	cui.ASCII = false

	//设计布局cui
	cui.SetManagerFunc(layout)
	// 注册回调事件
	//1.第一个是ctrl+c退出程序；
	if err := cui.SetKeybinding("main", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}
	if err := cui.SetKeybinding("main",gocui.KeyEnter, gocui.ModNone, viewUpdate); err != nil {
		log.Panicln(err)
	}
	if err := cui.SetKeybinding("main", gocui.KeyPgup, gocui.ModNone, viewUpScroll); err != nil {
		log.Panicln(err)
	}
	if err := cui.SetKeybinding("main", gocui.KeyPgdn, gocui.ModNone, viewDownScroll); err != nil {
		log.Panicln(err)
	}
	if err := cui.SetKeybinding("main", gocui.KeyArrowDown, gocui.ModNone, pasteDown); err != nil {
		log.Panicln(err)
	}
	if err := cui.SetKeybinding("main", gocui.KeyArrowUp, gocui.ModNone, pasteUP); err != nil {
		log.Panicln(err)
	}
	if err:=cui.MainLoop();err!=nil{
		log.Print(err)
	}
}

// 整体布局
func layout(g *gocui.Gui) error {
	maxX,maxY := g.Size()
	if err := headLayout(g, 1, 1, maxX-1, 3); err != nil {
		return err
	}
	if err:=outLayout(g,1,5,maxX-1,maxY-4);err!=nil{
		return err
	}
	if err:=mainLayout(g,1, maxY-3,maxX-1,maxY-1);err!=nil{
		return err
	}
	return nil
}

// 局部布局--聊天框的head部分
func headLayout(cui *gocui.Gui, x0, y0, x1, y1 int) error {
	if view, err := cui.SetView("head", x0, y0, x1, y1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		view.Wrap = false
		view.Overwrite = true
		fmt.Print("head头部分")
		msg := "开始聊天吧~"
		setHeadText(cui, msg)
	}
	return nil
}

//局部布局--main布局
func mainLayout(cui *gocui.Gui,x0,y0,x1,y1 int) error{
	if view,err:=cui.SetView("main",x0,y0,x1,y1);err!=nil{
		if err!=gocui.ErrUnknownView{
			return err
		}
		view.Editable=true
		view.Wrap=true
		view.Overwrite=false
		if _,err:=cui.SetCurrentView("main");err!=nil{
			return err
		}
	}
	return nil
}

//局部布局--out布局
func outLayout(cui *gocui.Gui,x0,y0,x1,y1 int) error{
	if view,err:=cui.SetView("out",x0,y0,x1,y1);err!=nil{
		if err!=gocui.ErrUnknownView{
			return err
		}
		view.Wrap=true
		view.Overwrite=false
		view.Autoscroll=true
		view.SelBgColor=gocui.ColorRed
		view.Title="Message"
	}
	return nil
}

/**
* 将聊天界面的头部填充内容
*/
func setHeadText(g *gocui.Gui, msg string) {
	view, err := g.View("head")
	if err == nil {
		view.Clear()
		fmt.Fprint(view, color.FgGreen.Text(msg))
	}
}

/**
* 绑定退出事件，
* 需要完成：①获取消息对象准备持久化 ②关闭连接
*/
func quit(cui *gocui.Gui,cv *gocui.View)error{
	chat.Close()
	ov,_:=cui.View("out")
	buf=ov.Buffer()
	cui.Close()
	return gocui.ErrQuit

}

/**
* 页面更新
*/
func viewUpdate(cui *gocui.Gui,cv *gocui.View) error {
	view,err:=cui.View("out")
	view.Autoscroll=false
	ox,oy:=view.Origin()
	if err==nil{
		view.SetOrigin(ox,oy-1)
	}
	return nil
}

/**
* 绑定是消息上下翻页，翻查上一页
*/
func viewUpScroll(cui *gocui.Gui,cv *gocui.View)error{
	view,err:=cui.View("out")
	view.Autoscroll=false
	ox,oy:=view.Origin()
	if err==nil{
		view.SetOrigin(ox,oy-1)
	}
	return nil
}


/**
* 绑定的是消息上下翻页，翻查下一页
*/ 
func viewDownScroll(cui *gocui.Gui,cv *gocui.View)error{
	view,err:=cui.View("out")
	_,y:=view.Size()
	ox,oy:=view.Origin()
	lnum:=len(view.BufferLines())
	if err==nil{
		if oy>lnum-y-1{
			view.Autoscroll=true
		}else{
			view.SetOrigin(ox,oy+1)
		}
	}
	return nil
}

/*
* 发送框查找上一条消息
*/
func pasteUP(cui *gocui.Gui,cv *gocui.View)error{
	view,err:=cui.View("out")
	if err!=nil{
		fmt.Fprint(cv,err)
		return nil
	}
	//获取到view的所有消息行,用来查找数据的上下行
	lines:=view.BufferLines()
	len:=len(lines)
	//判断是否有查找上一条消息的空间
	if pos<len-1 {
		pos++
	}
	cv.Clear()
	fmt.Fprintf(cv,"%s",lines[len-1-pos])
	return nil
}

/**
* 查找发送框的下一条消息，（后续看看能不能和pasteUP合成为一个）
*/
func pasteDown(cui *gocui.Gui,cv *gocui.View)error{
	view,err:=cui.View("out")
	if err!=nil{
		fmt.Fprint(cv,err)
		return nil
	}
	lines:=view.BufferLines()
	len:=len(lines)
	//判断是否有查找上一条消息的空间
	if pos>0 {
		pos--
	}
	cv.Clear()
	fmt.Fprintf(cv,"%s",lines[len-1-pos])
	return nil
}