// Code generated by qtc from "docs.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line templates/docs.qtpl:1
package templates

//line templates/docs.qtpl:1
import (
	"github.com/valyala/fasthttp"
)

//line templates/docs.qtpl:6
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line templates/docs.qtpl:6
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line templates/docs.qtpl:7
type DocsPage struct {
	CTX  *fasthttp.RequestCtx
	Js   string
	Css  string
	Host string
}

//line templates/docs.qtpl:15
func (view *DocsPage) StreamBody(qw422016 *qt422016.Writer) {
//line templates/docs.qtpl:15
	qw422016.N().S(`
<div style="margin:10px auto;width:700px;">

    <div style="margin:15px 0;">
        <div style="display: inline-block;vertical-align: top;width:50px;height:50px;border-radius:8px;border:1px solid #ccc; background-color: #fff;margin-right: 10px">

        </div>
        <div style="display: inline-block;vertical-align: top">
            <div style="font-weight: 600;font-size:18px;margin-top: 3px">URL Shortener</div>
            <div>Задача от Chaos за Николай Иванов</div>
        </div>
        <div style="float:right;display:flex;flex-direction: column">
            <a href="/task"> <div style="display: flex">
                <div style="background-color: #3F2D64;border-radius: 6px;width:25px;height:25px;line-height:  25px;   text-align: center;
 font-weight: bold;font-size:14px;color:#fff;">?
                </div>
                <div style="padding:3px 0 0 5px;">Условие на задачата</div>


            </div></a>
            <a href="/tech-doc"><div style="display: flex;margin-top: 5px">
                <div style="background-color: #3F2D64;border-radius: 6px;width:25px;height:25px;line-height:  23px;   text-align: center;
 font-weight: bold;font-size:14px;color:#fff;">{&nbsp;}
                </div>
                <div style="padding:3px 0 0 5px;">Техническа документация</div>


            </div></a>
        </div>
    </div>
    <div class="method"
         style="background-color: #fff;border-radius:6px;overflow:auto;padding:8px;border:1px solid #ccc;border-bottom:2px solid #ccc;margin:5px 0 10px 0;">
        <div class="action" style="border: 1px solid #D2D8E4;
    background-color: #F6F7FA;
    border-radius: 4px;
    color: #3A72F9;
    padding: 5px 0;
    font-weight: bold;
    display: inline-block;
    width: 55px;
    text-align: center;">
            POST
        </div>
        <div class="route" style="    border: 1px solid #B9CDFD;
    background-color: #fff;
    border-radius: 4px;
    color: #303562;
    padding: 5px 10px;
    font-weight: 600;
    display: inline-block;
    margin-left: 5px;">
            /
        </div>
        <div class="tabs">
            <div class="tab on" data-tab="description">Описание</div>
            <div class="tab" data-tab="parameters">Параметри</div>
            <div class="tab" data-tab="test">Тестване</div>
        </div>
        <div class="sections">
            <div class="tab-section on" data-tab="description">Връща кратък адрес след постване на пълния адрес.</div>
            <div class="tab-section" data-tab="parameters"></div>
            <div class="tab-section" data-tab="test">
                <div style="display: flex">
                    <div><input class="param" type="text" placeholder="Въведете адресът, който искате да скъсите" style="width:400px;"></div>
                    <button class="post-test-button">Изпрати</button>
                </div>
                <div class="result"></div>
            </div>
        </div>
    </div>
    <div class="method"
         style="background-color: #fff;border-radius:6px;overflow:auto;padding:8px;border:1px solid #ccc;border-bottom:2px solid #ccc;margin:5px 0 10px 0;">
        <div class="action" style="border: 1px solid #D2D8E4;
    background-color: #F6F7FA;
    border-radius: 4px;
    color: #3A72F9;
    padding: 5px 0;
    font-weight: bold;
    display: inline-block;
    width: 55px;
    text-align: center;">
            GET
        </div>
        <div class="route" style="    border: 1px solid #B9CDFD;
    background-color: #fff;
    border-radius: 4px;
    color: #303562;
    padding: 5px 10px;
    font-weight: 600;
    display: inline-block;
    margin-left: 5px;">
            /{<span style="color:#AB4C7F;">кратък-код</span>}
        </div>
        <div class="tabs">
            <div class="tab on" data-tab="description">Описание</div>
            <div class="tab" data-tab="parameters">Параметри</div>
            <div class="tab" data-tab="test">Тестване</div>
        </div>
        <div class="sections">
            <div class="tab-section on" data-tab="description">Пренасочва към оригиналния адрес.</div>
            <div class="tab-section" data-tab="parameters"></div>
            <div class="tab-section" data-tab="test"></div>
        </div>
    </div>
    <div style="background-color: #fff;  padding:10px;  border: 1px solid #ccc;
    border-bottom: 2px solid #ccc;">
        <div style="margin:0 0 5px 0;font-weight:600;">Използвани технологии:</div>
        <div>
<div style="padding:5px 0;">Програмен език: <span style="font-weight:600;">Go</span></div>
            <div style="padding:5px 0;">Хранилище: <span style="font-weight:600;">Couchbase</span></div>
            <div style="padding:5px 0;">Избрал съм тази система за базаданни, защото изискваме скорост при четене. Когато хората заявят кратък адрес, четенето от базата данни трябва да е много бързо. Couchbase ни предоставя най-добрата скорост на четене от базата данни, които поддържат както in-memory, така и persistent disk storage.
            Другата причина е естеството на данните. Трябва ни Key/Value хранилище, където ключът е краткия, а стойността е оригиналния адрес.
            </div>
            <div style="padding:5px 0;">Използвани библиотеки:
                <ul style=" margin: 10px 0 10px 20px;">
                    <li><span style="font-weight:600;">fasthttp/router</span> - най-бързата http библиотека за Go</li>
                    <li><span style="font-weight:600;">quicktemplate</span> - най-бързата template библиотека за Go</li>
                </ul>
            </div>
        </div>
    </div>
    <div style="background-color: #fff; margin-top: 10px; padding:10px;  border: 1px solid #ccc;
    border-bottom: 2px solid #ccc;">
        <div style="margin:0 0 5px 0;font-weight:600;">Забележка към условието на задачата:</div>
        <div>Две от функционалните изисквания на задачата си противоречат. От една страна задачата изисква размера на краткия адрес да е не повече от 5 символа, а от друга да може да побира 1 милиард записа.
            Това е невъзможно тъй като възможните комбинации (пермутации) на разрешените символи е 8,936,928. В реален продукт краткият адрес трябва да е най-малко 7 символа, за да могат да се генерират 1 милиард уникални хеша.</div>
        <div>
            <div >
                <ul style=" margin: 10px 0 10px 20px;">
                    <li>Възможни символи (английски главни [26] и малки букви [26] + числата [10]): <span style="font-weight:600;">62</span></li>
                    <li>Максимален брой символи: <span style="font-weight:600;">5</span></li>
                    <li>Брой уникални комбинация (включващи повторенията на символите): <span style="font-weight:600;">8,936,928</span></li>
                </ul>
            </div>
            <img src="/gfx/permutations.png">
        </div>
    </div>

</div>
`)
//line templates/docs.qtpl:154
}

//line templates/docs.qtpl:154
func (view *DocsPage) WriteBody(qq422016 qtio422016.Writer) {
//line templates/docs.qtpl:154
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/docs.qtpl:154
	view.StreamBody(qw422016)
//line templates/docs.qtpl:154
	qt422016.ReleaseWriter(qw422016)
//line templates/docs.qtpl:154
}

//line templates/docs.qtpl:154
func (view *DocsPage) Body() string {
//line templates/docs.qtpl:154
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/docs.qtpl:154
	view.WriteBody(qb422016)
//line templates/docs.qtpl:154
	qs422016 := string(qb422016.B)
//line templates/docs.qtpl:154
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/docs.qtpl:154
	return qs422016
//line templates/docs.qtpl:154
}

//line templates/docs.qtpl:155
func (view *DocsPage) StreamDescription(qw422016 *qt422016.Writer) {
//line templates/docs.qtpl:155
}

//line templates/docs.qtpl:155
func (view *DocsPage) WriteDescription(qq422016 qtio422016.Writer) {
//line templates/docs.qtpl:155
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/docs.qtpl:155
	view.StreamDescription(qw422016)
//line templates/docs.qtpl:155
	qt422016.ReleaseWriter(qw422016)
//line templates/docs.qtpl:155
}

//line templates/docs.qtpl:155
func (view *DocsPage) Description() string {
//line templates/docs.qtpl:155
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/docs.qtpl:155
	view.WriteDescription(qb422016)
//line templates/docs.qtpl:155
	qs422016 := string(qb422016.B)
//line templates/docs.qtpl:155
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/docs.qtpl:155
	return qs422016
//line templates/docs.qtpl:155
}

//line templates/docs.qtpl:156
func (view *DocsPage) StreamTitle(qw422016 *qt422016.Writer) {
//line templates/docs.qtpl:156
}

//line templates/docs.qtpl:156
func (view *DocsPage) WriteTitle(qq422016 qtio422016.Writer) {
//line templates/docs.qtpl:156
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/docs.qtpl:156
	view.StreamTitle(qw422016)
//line templates/docs.qtpl:156
	qt422016.ReleaseWriter(qw422016)
//line templates/docs.qtpl:156
}

//line templates/docs.qtpl:156
func (view *DocsPage) Title() string {
//line templates/docs.qtpl:156
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/docs.qtpl:156
	view.WriteTitle(qb422016)
//line templates/docs.qtpl:156
	qs422016 := string(qb422016.B)
//line templates/docs.qtpl:156
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/docs.qtpl:156
	return qs422016
//line templates/docs.qtpl:156
}

//line templates/docs.qtpl:157
func (view *DocsPage) StreamDomain(qw422016 *qt422016.Writer) {
//line templates/docs.qtpl:157
	qw422016.E().S(view.Host)
//line templates/docs.qtpl:157
}

//line templates/docs.qtpl:157
func (view *DocsPage) WriteDomain(qq422016 qtio422016.Writer) {
//line templates/docs.qtpl:157
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/docs.qtpl:157
	view.StreamDomain(qw422016)
//line templates/docs.qtpl:157
	qt422016.ReleaseWriter(qw422016)
//line templates/docs.qtpl:157
}

//line templates/docs.qtpl:157
func (view *DocsPage) Domain() string {
//line templates/docs.qtpl:157
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/docs.qtpl:157
	view.WriteDomain(qb422016)
//line templates/docs.qtpl:157
	qs422016 := string(qb422016.B)
//line templates/docs.qtpl:157
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/docs.qtpl:157
	return qs422016
//line templates/docs.qtpl:157
}

//line templates/docs.qtpl:158
func (view *DocsPage) StreamStylesheet(qw422016 *qt422016.Writer) {
//line templates/docs.qtpl:158
	qw422016.E().S(view.Css)
//line templates/docs.qtpl:158
}

//line templates/docs.qtpl:158
func (view *DocsPage) WriteStylesheet(qq422016 qtio422016.Writer) {
//line templates/docs.qtpl:158
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/docs.qtpl:158
	view.StreamStylesheet(qw422016)
//line templates/docs.qtpl:158
	qt422016.ReleaseWriter(qw422016)
//line templates/docs.qtpl:158
}

//line templates/docs.qtpl:158
func (view *DocsPage) Stylesheet() string {
//line templates/docs.qtpl:158
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/docs.qtpl:158
	view.WriteStylesheet(qb422016)
//line templates/docs.qtpl:158
	qs422016 := string(qb422016.B)
//line templates/docs.qtpl:158
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/docs.qtpl:158
	return qs422016
//line templates/docs.qtpl:158
}

//line templates/docs.qtpl:159
func (view *DocsPage) StreamScript(qw422016 *qt422016.Writer) {
//line templates/docs.qtpl:159
	qw422016.E().S(view.Js)
//line templates/docs.qtpl:159
}

//line templates/docs.qtpl:159
func (view *DocsPage) WriteScript(qq422016 qtio422016.Writer) {
//line templates/docs.qtpl:159
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/docs.qtpl:159
	view.StreamScript(qw422016)
//line templates/docs.qtpl:159
	qt422016.ReleaseWriter(qw422016)
//line templates/docs.qtpl:159
}

//line templates/docs.qtpl:159
func (view *DocsPage) Script() string {
//line templates/docs.qtpl:159
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/docs.qtpl:159
	view.WriteScript(qb422016)
//line templates/docs.qtpl:159
	qs422016 := string(qb422016.B)
//line templates/docs.qtpl:159
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/docs.qtpl:159
	return qs422016
//line templates/docs.qtpl:159
}