package foaas

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const endpoint string = "https://foaas.com"

type Client struct {
	Lang   string
	Format string
	Shout  bool
}

func makeRequest(operation string, cl *Client) (string, error) {
	client := http.Client{}
	req, err := http.NewRequest(http.MethodGet, endpoint+operation, nil)
	if err != nil {
		return "", nil
	}

	var hval string
	switch cl.Format {
	case "json":
		hval = "application/json"
		break
	case "xml":
		hval = "application/xml"
		break
	case "html":
		hval = "text/html"
		break
	default:
		hval = "text/plain"
	}
	req.Header.Set("Accept", hval)

	qry := req.URL.Query()
	if cl.Lang != "" {
		qry.Add("i18n", cl.Lang)
	}
	if cl.Shout == true {
		qry.Add("shoutcloud", "")
	}
	req.URL.RawQuery = qry.Encode()

	res, err := client.Do(req)
	if err != nil {
		return "", nil
	}
	defer res.Body.Close()

	respBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", nil
	}

	return string(respBody), nil
}

/*	Version
	Will return content with the current FOAAS version number.
*/
func (cl *Client) Version() (string, error) {
	return makeRequest("/version", cl)
}

/*	Operations
	Will return a JSON list of operations with names and fields. Note: JSON Only.
*/
func (cl *Client) Operations() (string, error) {
	return makeRequest("/operations", cl)
}

/*	Anyway
	Will return content of the form "Who the fuck are you anyway, <company>,
	why are you stirring up so much trouble, and, who pays you? - <from>"
*/
func (cl *Client) Anyway(company string, from string) (string, error) {
	url := fmt.Sprintf("/anyway/%s/%s", company, from)
	return makeRequest(url, cl)
}

/*	Asshole
	Will return content of the form "Fuck you, asshole. - <from>"
*/
func (cl *Client) Asshole(from string) (string, error) {
	url := fmt.Sprintf("/asshole/%s", from)
	return makeRequest(url, cl)
}

/*	Awesome
	Will return content of the form "This is Fucking Awesome. - <from>"
*/
func (cl *Client) Awesome(from string) (string, error) {
	url := fmt.Sprintf("/awesome/%s", from)
	return makeRequest(url, cl)
}

/*	Back
	Will return content of the form "<name>, back the fuck off. - <from>"
*/
func (cl *Client) Back(name string, from string) (string, error) {
	url := fmt.Sprintf("/back/%s/%s", name, from)
	return makeRequest(url, cl)
}

/*	Bag
	Will return content of the form "Eat a bag of fucking dicks. - <from>"
*/
func (cl *Client) Bag(from string) (string, error) {
	url := fmt.Sprintf("/bag/%s", from)
	return makeRequest(url, cl)
}

/*	Ballmer
	Will return content of the form "Fucking <name> is a fucking pussy.
	I'm going to fucking bury that guy, I have done it before, and I will do it again.
	I'm going to fucking kill <company>. - <from>"
*/
func (cl *Client) Ballmer(name string, company string, from string) (string, error) {
	url := fmt.Sprintf("/ballmer/%s/%s/%s", name, company, from)
	return makeRequest(url, cl)
}

/*	Birthday
	Will return content of the form "Happy Fucking Birthday, <name>. - <from>"
*/
func (cl *Client) Birthday(name string, from string) (string, error) {
	url := fmt.Sprintf("/bday/%s/%s", name, from)
	return makeRequest(url, cl)
}

/*	Because
	Will return content of the form "Why? Because fuck you, that's why. - <from>"
*/
func (cl *Client) Because(from string) (string, error) {
	url := fmt.Sprintf("/because/%s", from)
	return makeRequest(url, cl)
}

/*	Blackadder
	Will return content of the form
	"<name>, your head is as empty as a eunuch’s underpants. Fuck off! - <from>"
*/
func (cl *Client) Blackadder(name string, from string) (string, error) {
	url := fmt.Sprintf("/blackadder/%s/%s", name, from)
	return makeRequest(url, cl)
}

/*	BravoMike
	Will return content of the form "Bravo mike, <name>. - <from>"
*/
func (cl *Client) BravoMike(name string, from string) (string, error) {
	url := fmt.Sprintf("/bm/%s/%s", name, from)
	return makeRequest(url, cl)
}

/*	Bucket
	Will return content of the form "Please choke on a bucket of cocks. - <from>"
*/
func (cl *Client) Bucket(from string) (string, error) {
	url := fmt.Sprintf("/bucket/%s", from)
	return makeRequest(url, cl)
}

/*	Bus
	Will return content of the form "Christ on a bendy-bus, <name>,
	don't be such a fucking faff-arse. - <from>"
*/
func (cl *Client) Bus(name string, from string) (string, error) {
	url := fmt.Sprintf("/bus/%s/%s", name, from)
	return makeRequest(url, cl)
}

/*	Bye
	Will return content of the form <Fuckity bye! - <from>"
*/
func (cl *Client) Bye(from string) (string, error) {
	url := fmt.Sprintf("/bye/%s", from)
	return makeRequest(url, cl)
}

/*	CanIUse
	Will return content of the form "Can you use <tool>? Fuck no! - <from>"
*/
func (cl *Client) CanIUse(tool string, from string) (string, error) {
	url := fmt.Sprintf("/caniuse/%s/%s", tool, from)
	return makeRequest(url, cl)
}

/*	Chainsaw
	Will return content of the form
	"Fuck me gently with a chainsaw, <name>. Do I look like Mother Teresa? - <from>"
*/
func (cl *Client) Chainsaw(name string, from string) (string, error) {
	url := fmt.Sprintf("/chainsaw/%s/%s", name, from)
	return makeRequest(url, cl)
}

/*	Cocksplat
	Will return content of the form "Fuck off <name>, you worthless cocksplat - <from>"
*/
func (cl *Client) Cocksplat(name string, from string) (string, error) {
	url := fmt.Sprintf("/cocksplat/%s/%s", name, from)
	return makeRequest(url, cl)
}

/*	Cool
	Will return content of the form "Cool story, bro. - <from>"
*/
func (cl *Client) Cool(from string) (string, error) {
	url := fmt.Sprintf("/cool/%s", from)
	return makeRequest(url, cl)
}

/*	Cup
	Will return content of the form "How about a nice cup of shut the fuck up? - <from>"
*/
func (cl *Client) Cup(from string) (string, error) {
	url := fmt.Sprintf("/cup/%s", from)
	return makeRequest(url, cl)
}

/*	Dalton
	Will return content of the form "<name> A fucking problem solving super-hero. - <from>"
*/
func (cl *Client) Dalton(name string, from string) (string, error) {
	url := fmt.Sprintf("/dalton/%s/%s", name, from)
	return makeRequest(url, cl)
}

/*	Deraadt
	Will return content of the form
	"<name> you are being the usual slimy hypocritical asshole...
	You may have had value ten years ago,
	but people will see that you don't anymore. - <from>"
*/
func (cl *Client) Deraadt(name string, from string) (string, error) {
	url := fmt.Sprintf("/deraadt/%s/%s", name, from)
	return makeRequest(url, cl)
}

/*	Diabetes
	Will return content of the form
	"I'd love to stop and chat to you but I'd rather have type 2 diabetes. - <from>"
*/
func (cl *Client) Diabetes(from string) (string, error) {
	url := fmt.Sprintf("/diabetes/%s", from)
	return makeRequest(url, cl)
}

/*	Donut
	Will return content of the form "<name>, go and take a flying fuck at a rolling donut. - <from>"
*/
func (cl *Client) Donut(name string, from string) (string, error) {
	url := fmt.Sprintf("/donut/%s/%s", name, from)
	return makeRequest(url, cl)
}

/*	DoSomething
	Will return content of the form "<do> the fucking <something>! - <from>"
*/
func (cl *Client) DoSomething(do string, something string, from string) (string, error) {
	url := fmt.Sprintf("/dosomething/%s/%s/%s", do, something, from)
	return makeRequest(url, cl)
}

/*	Equity
	Will return content of the form
	"Equity only? Long hours? Zero Pay? Well <name>, just sign me right the fuck up. - <from>"
*/
func (cl *Client) Equity(name string, from string) (string, error) {
	url := fmt.Sprintf("/equity/%s/%s", name, from)
	return makeRequest(url, cl)
}

/*	Even
	Will return content of the form "I can't fuckin' even. - <from>"
*/
func (cl *Client) Even(from string) (string, error) {
	url := fmt.Sprintf("/even/%s", from)
	return makeRequest(url, cl)
}

/*	Everyone
	Will return content of the form "Everyone can go and fuck off. - <from>"
*/
func (cl *Client) Everyone(from string) (string, error) {
	url := fmt.Sprintf("/everyone/%s", from)
	return makeRequest(url, cl)
}

/*	Everything
	Will return content of the form "Fuck everything. - <from>"
*/
func (cl *Client) Everything(from string) (string, error) {
	url := fmt.Sprintf("/everything/%s", from)
	return makeRequest(url, cl)
}

/*	Family
	Will return content of the form "Fuck you, your whole family, your pets, and your feces. - <from>"
*/
func (cl *Client) Family(from string) (string, error) {
	url := fmt.Sprintf("/family/%s", from)
	return makeRequest(url, cl)
}

/*	Fascinating
	Will return content of the form "Fascinating story, in what chapter do you shut the fuck up? - <from>"
*/
func (cl *Client) Fascinating(from string) (string, error) {
	url := fmt.Sprintf("/fascinating/%s", from)
	return makeRequest(url, cl)
}

/*	Fewer
	Will return content of the form "Go fuck yourself <name>, you'll disappoint fewer people. - <from>"
*/
func (cl *Client) Fewer(name string, from string) (string, error) {
	url := fmt.Sprintf("/fewer/%s/%s", name, from)
	return makeRequest(url, cl)
}

/*	Field
	Will return content of the form
	"And <name> said unto <from>, 'Verily, cast thine eyes upon the field in which I grow my fucks',
	and <from> gave witness unto the field, and saw that it was barren. - <reference>"
*/
func (cl *Client) Field(name string, from string, reference string) (string, error) {
	url := fmt.Sprintf("/field/%s/%s/%s", name, from, reference)
	return makeRequest(url, cl)
}

/*	Flying
	Will return content of the form "I don't give a flying fuck. - <from>"
*/
func (cl *Client) Flying(from string) (string, error) {
	url := fmt.Sprintf("/flying/%s", from)
	return makeRequest(url, cl)
}

/*	Ftfy
	Will return content of the form "Fuck That, Fuck You - <from>"
*/
func (cl *Client) Ftfy(from string) (string, error) {
	url := fmt.Sprintf("/ftfy/%s", from)
	return makeRequest(url, cl)
}

/*	Fts
	Will return content of the form "Fuck that shit, <name>. - <from>"
*/
func (cl *Client) Fts(name string, from string) (string, error) {
	url := fmt.Sprintf("/fts/%s/%s", name, from)
	return makeRequest(url, cl)
}

/*	Fyyff
	Will return content of the form "Fuck you, you fucking fuck. - "from>"
*/
func (cl *Client) Fyyff(from string) (string, error) {
	url := fmt.Sprintf("/fyyff/%s", from)
	return makeRequest(url, cl)
}

/*	Gfy
	Will return content of the form "Golf foxtrot yankee, <name>. - <from>"
*/
func (cl *Client) Gfy(name string, from string) (string, error) {
	url := fmt.Sprintf("/gfy/%s/%s", name, from)
	return makeRequest(url, cl)
}

/*	Give
	Will return content of the form "I give zero fucks. - <from>"
*/
func (cl *Client) Give(from string) (string, error) {
	url := fmt.Sprintf("/give/%s", from)
	return makeRequest(url, cl)
}

/*	Greed
	Will return content of the form
	"The point is, ladies and gentleman, that <noun> -- for lack of a better word -- is good.
	<noun> is right. <noun> works. <noun> clarifies, cuts through,
	and captures the essence of the evolutionary spirit.
	<noun>, in all of its forms -- <noun> for life, for money, for love,
	knowledge -- has marked the upward surge of mankind. - <from>"
*/
func (cl *Client) Greed(noun string, from string) (string, error) {
	url := fmt.Sprintf("/greed/%s/%s", noun, from)
	return makeRequest(url, cl)
}

/*	HolyGrail
	Will return content of the form
	"I don't want to talk to you, no more, you empty-headed animal, food trough wiper.
	I fart in your general direction. Your mother was a hamster and your father smelt of elderberries.
	Now go away or I shall taunt you a second time. - <from>"
*/
func (cl *Client) HolyGrail(from string) (string, error) {
	url := fmt.Sprintf("/holygrail/%s", from)
	return makeRequest(url, cl)
}

/*	Horse
	Will return content of the form "Fuck you and the horse you rode in on. - <from>"
*/
func (cl *Client) Horse(from string) (string, error) {
	url := fmt.Sprintf("/horse/%s", from)
	return makeRequest(url, cl)
}

/*	Idea
	Will return content of the form "That sounds like a fucking great idea! - <from>"
*/
func (cl *Client) Idea(from string) (string, error) {
	url := fmt.Sprintf("/idea/%s", from)
	return makeRequest(url, cl)
}

/*	Immensity
	Will return content of the form "You can not imagine the immensity of the FUCK I do not give. - <from>"
*/
func (cl *Client) Immensity(from string) (string, error) {
	url := fmt.Sprintf("/immensity/%s", from)
	return makeRequest(url, cl)
}

/*	Ing
	Will return content of the form "Fucking fuck off, <name>. - <from>"
*/
func (cl *Client) Ing(name string, from string) (string, error) {
	url := fmt.Sprintf("/ing/%s/%s", name, from)
	return makeRequest(url, cl)
}

/*	JingleBells
	Will return content of the form
	"Fuck you, fuck me, fuck your family. Fuck your father, fuck your mother, fuck you and me. - <from>"
*/
func (cl *Client) JingleBells(from string) (string, error) {
	url := fmt.Sprintf("/jinglebells/%s", from)
	return makeRequest(url, cl)
}

/*	Keep
	Will return content of the form
	"<name> Fuck off. And when you get there, fuck off from there too. Then fuck off some more.
	Keep fucking off until you get back here. Then fuck off again. - <from>"
*/
func (cl *Client) Keep(name string, from string) (string, error) {
	url := fmt.Sprintf("/keep/%s/%s", name, from)
	return makeRequest(url, cl)
}

/*	KeepCalm
	Will return content of the form "Keep the fuck calm and <reaction>! - <from>"
*/
func (cl *Client) KeepCalm(reaction string, from string) (string, error) {
	url := fmt.Sprintf("/keepcalm/%s/%s", reaction, from)
	return makeRequest(url, cl)
}

/*	King
	Will return content of the form "Oh fuck off, just really fuck off you total dickface.
	Christ, <name>, you are fucking thick. - <from>"
*/
func (cl *Client) King(name string, from string) (string, error) {
	url := fmt.Sprintf("/king/%s/%s", name, from)
	return makeRequest(url, cl)
}

/*	Legend
	Will return content of the form "<name>, you're a fucking legend. - <from>"
*/
func (cl *Client) Legend(name string, from string) (string, error) {
	url := fmt.Sprintf("/legend/%s/%s", name, from)
	return makeRequest(url, cl)
}

/*	Life
	Will return content of the form "Fuck my life. - <from>"
*/
func (cl *Client) Life(from string) (string, error) {
	url := fmt.Sprintf("/life/%s", from)
	return makeRequest(url, cl)
}

/*	Linus
	Will return content of the form
	"<name>, there aren't enough swear-words in the English language,
	so now I'll have to call you perkeleen vittupää just to express my
	disgust and frustration with this crap. - <from>"
*/
func (cl *Client) Linus(from string) (string, error) {
	url := fmt.Sprintf("/linus/%s", from)
	return makeRequest(url, cl)
}

/*	Logs
	Will return content of the form "Check your fucking logs! - <from>"
*/
func (cl *Client) Logs(from string) (string, error) {
	url := fmt.Sprintf("/logs/%s", from)
	return makeRequest(url, cl)
}

/*	Look
	Will return content of the form "<name>, do I look like I give a fuck? - <from>"
*/
func (cl *Client) Look(name string, from string) (string, error) {
	url := fmt.Sprintf("/look/%s/%s", name, from)
	return makeRequest(url, cl)
}

/*	Looking
	Will return content of the form "Looking for a fuck to give. - <from>"
*/
func (cl *Client) Looking(from string) (string, error) {
	url := fmt.Sprintf("/looking/%s", from)
	return makeRequest(url, cl)
}

/*	Madison
	Will return content of the form
	"What you've just said is one of the most insanely idiotic things I have ever heard, <name>.
	At no point in your rambling, incoherent response were you even close to anything that could
	be considered a rational thought. Everyone in this room is now dumber for having listened to it.
	I award you no points <name>, and may God have mercy on your soul. - <from>"
*/
func (cl *Client) Madison(name string, from string) (string, error) {
	url := fmt.Sprintf("/madison/%s/%s", name, from)
	return makeRequest(url, cl)
}

/*	Maybe
	Will return content of the form "Maybe. Maybe not. Maybe fuck yourself. - <from>"
*/
func (cl *Client) Maybe(from string) (string, error) {
	url := fmt.Sprintf("/Maybe/%s", from)
	return makeRequest(url, cl)
}

/*	Me
	Will return content of the form "Fuck me. - <from>"
*/
func (cl *Client) Me(from string) (string, error) {
	url := fmt.Sprintf("/me/%s", from)
	return makeRequest(url, cl)
}

/*	Morning
	Will return content of the form "Happy fuckin' mornin'! - <from>"
*/
func (cl *Client) Morning(from string) (string, error) {
	url := fmt.Sprintf("/mornin/%s", from)
	return makeRequest(url, cl)
}

/*	No
	Will return content of the form "No fucks given. - <from>"
*/
func (cl *Client) No(from string) (string, error) {
	url := fmt.Sprintf("/no/%s", from)
	return makeRequest(url, cl)
}

/*	Nugget
	Will return content of the form "Well <name>, aren't you a shining example of a rancid fuck-nugget. - <from>"
*/
func (cl *Client) Nugget(name string, from string) (string, error) {
	url := fmt.Sprintf("/nugget/%s%s", name, from)
	return makeRequest(url, cl)
}

/*	Off
	Will return content of the form "Fuck off, <name>. - <from>"
*/
func (cl *Client) Off(name string, from string) (string, error) {
	url := fmt.Sprintf("/off/%s/%s", name, from)
	return makeRequest(url, cl)
}

/*	OffWith
	Will return content of the form "Fuck off with <behavior> - <from>"
*/
func (cl *Client) OffWith(behavior string, from string) (string, error) {
	url := fmt.Sprintf("/off-with/%s/%s", behavior, from)
	return makeRequest(url, cl)
}

/*	Outside
	Will return content of the form "<name>, why don't you go outside and play hide-and-go-fuck-yourself? - <from>"
*/
func (cl *Client) Outside(name string, from string) (string, error) {
	url := fmt.Sprintf("/outside/%s/%s", name, from)
	return makeRequest(url, cl)
}

/*	Particular
	Will return content of the form "Fuck this <thing> in particular. - <from>"
*/
func (cl *Client) Particular(thing string, from string) (string, error) {
	url := fmt.Sprintf("/particular/%s/%s", thing, from)
	return makeRequest(url, cl)
}

/*	Pink
	Will return content of the form "Well, fuck me pink. - <from>"
*/
func (cl *Client) Pink(from string) (string, error) {
	url := fmt.Sprintf("/pink/%s", from)
	return makeRequest(url, cl)
}

/*	Problem
	Will return content of the form "What the fuck is your problem <name>? - <from>"
*/
func (cl *Client) Problem(name string, from string) (string, error) {
	url := fmt.Sprintf("/problem/%s/%s", name, from)
	return makeRequest(url, cl)
}

/*	Programmer
	Will return content of the form "Fuck you, I'm a programmer, bitch! - <from>"
*/
func (cl *Client) Programmer(from string) (string, error) {
	url := fmt.Sprintf("/programmer/%s", from)
	return makeRequest(url, cl)
}

/*	Pulp
	Will return content of the form "<language>, motherfucker, do you speak it? - <from>"
*/
func (cl *Client) Pulp(language string, from string) (string, error) {
	url := fmt.Sprintf("/pulp/%s/%s", language, from)
	return makeRequest(url, cl)
}

/*	Questions
	Will return content of the form "To fuck off, or to fuck off (that is not a question) - <from>"
*/
func (cl *Client) Question(from string) (string, error) {
	url := fmt.Sprintf("/question/%s", from)
	return makeRequest(url, cl)
}

/*	RatsArse
	Will return content of the form "I don't give a rat's arse. - <from>"
*/
func (cl *Client) RatsArse(from string) (string, error) {
	url := fmt.Sprintf("/typehere/%s", from)
	return makeRequest(url, cl)
}

/*	Retard
	Will return content of the form "You Fucktard! - <from>"
*/
func (cl *Client) Retard(from string) (string, error) {
	url := fmt.Sprintf("/retard/%s", from)
	return makeRequest(url, cl)
}

/*	Ridiculous
	Will return content of the form "That's fucking ridiculous - <from>"
*/
func (cl *Client) Ridiculous(from string) (string, error) {
	url := fmt.Sprintf("/ridiculous/%s", from)
	return makeRequest(url, cl)
}

/*	Rockstar
	Will return content of the form "<name>, you're a fucking Rock Star! - <from>"
*/
func (cl *Client) Rockstar(name string, from string) (string, error) {
	url := fmt.Sprintf("/rockstar/%s/%s", name, from)
	return makeRequest(url, cl)
}

/*	Rtfm
	Will return content of the form "Read the fucking manual! - <from>"
*/
func (cl *Client) Rtfm(from string) (string, error) {
	url := fmt.Sprintf("/rtfm/%s", from)
	return makeRequest(url, cl)
}

/*	Sake
	Will return content of the form "For Fuck's sake! - <from>"
*/
func (cl *Client) Sake(from string) (string, error) {
	url := fmt.Sprintf("/sake/%s", from)
	return makeRequest(url, cl)
}

/*	Shakespeare
	Will return content of the form
	"<name>, Thou clay-brained guts, thou knotty-pated fool,
	thou whoreson obscene greasy tallow-catch! - <from>"
*/
func (cl *Client) Shakespeare(name string, from string) (string, error) {
	url := fmt.Sprintf("/shakespeare/%s/%s", name, from)
	return makeRequest(url, cl)
}

/*	Shit
	Will return content of the form "Fuck this shit! - <from>"
*/
func (cl *Client) Shit(from string) (string, error) {
	url := fmt.Sprintf("/shit/%s", from)
	return makeRequest(url, cl)
}

/*	ShutUp
	Will return content of the form "<name>, shut the fuck up. - <from>"
*/
func (cl *Client) ShutUp(name string, from string) (string, error) {
	url := fmt.Sprintf("/shutup/%s/%s", name, from)
	return makeRequest(url, cl)
}

/*	Single
	Will return content of the form "Not a single fuck was given. - <from>"
*/
func (cl *Client) Single(from string) (string, error) {
	url := fmt.Sprintf("/single/%s", from)
	return makeRequest(url, cl)
}

/*	Thanks
	Will return content of the form "Fuck you very much. - <from>"
*/
func (cl *Client) Thanks(from string) (string, error) {
	url := fmt.Sprintf("/thanks/%s", from)
	return makeRequest(url, cl)
}

/*	That
	Will return content of the form "Fuck that. - <from>"
*/
func (cl *Client) That(from string) (string, error) {
	url := fmt.Sprintf("/that/%s", from)
	return makeRequest(url, cl)
}

/*	Think
	Will return content of the form "<name>, you think I give a fuck? - <from>"
*/
func (cl *Client) Think(name string, from string) (string, error) {
	url := fmt.Sprintf("/think/%s/%s", name, from)
	return makeRequest(url, cl)
}

/*	Thinking
	Will return content of the form "<name>, what the fuck were you actually thinking? - <from>"
*/
func (cl *Client) Thinking(name string, from string) (string, error) {
	url := fmt.Sprintf("/thinking/%s/%s", name, from)
	return makeRequest(url, cl)
}

/*	This
	Will return content of the form "Who has two thumbs and doesn't give a fuck? <name>. - <from>"
*/
func (cl *Client) This(from string) (string, error) {
	url := fmt.Sprintf("/this/%s", from)
	return makeRequest(url, cl)
}

/*	Thumbs
	Will return content of the form "Fuck this. - <from>"
*/
func (cl *Client) Thumbs(name string, from string) (string, error) {
	url := fmt.Sprintf("/thumbs/%s/%s", name, from)
	return makeRequest(url, cl)
}

/*	Too
	Will return content of the form "Thanks, fuck you too. - <from>"
*/
func (cl *Client) Too(from string) (string, error) {
	url := fmt.Sprintf("/too/%s", from)
	return makeRequest(url, cl)
}

/*	Tucker
	Will return content of the form "Come the fuck in or fuck the fuck off. - <from>"
*/
func (cl *Client) Tucker(from string) (string, error) {
	url := fmt.Sprintf("/tucker/%s", from)
	return makeRequest(url, cl)
}

/*	Waste
	Will return content of the form "I don't waste my fucking time with your bullshit <name>! - <from>"
*/
func (cl *Client) Waste(name string, from string) (string, error) {
	url := fmt.Sprintf("/waste/%s/%s", name, from)
	return makeRequest(url, cl)
}

/*	What
	Will return content of the form "What the fuck‽ - <from>"
*/
func (cl *Client) What(from string) (string, error) {
	url := fmt.Sprintf("/what/%s", from)
	return makeRequest(url, cl)
}

/*	Xmas
	Will return content of the form "Merry Fucking Christmas, <name>. - <from>"
*/
func (cl *Client) Xmas(name string, from string) (string, error) {
	url := fmt.Sprintf("/xmas/%s/%s", name, from)
	return makeRequest(url, cl)
}

/*	Yoda
	Will return content of the form "Fuck off, you must, <name>. - <from>"
*/
func (cl *Client) Yoda(name string, from string) (string, error) {
	url := fmt.Sprintf("/yoda/%s/%s", name, from)
	return makeRequest(url, cl)
}

/*	You
	Will return content of the form "Fuck you, <name>. - <from>"
*/
func (cl *Client) You(name string, from string) (string, error) {
	url := fmt.Sprintf("/you/%s/%s", name, from)
	return makeRequest(url, cl)
}

/*	Zayn
	Will return content of the form "Ask me if I give a motherfuck ?!! - <from>"
*/
func (cl *Client) Zayn(from string) (string, error) {
	url := fmt.Sprintf("/zayn/%s", from)
	return makeRequest(url, cl)
}

/*	Zero
	Will return content of the form "Zero, that's the number of fucks I give. - <from>"
*/
func (cl *Client) Zero(from string) (string, error) {
	url := fmt.Sprintf("/zero/%s", from)
	return makeRequest(url, cl)
}
