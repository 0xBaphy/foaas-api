package foaas_api

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const endpoint string = "https://foaas.com"

type Foaas struct {
	lang   string
	format string
	shout  bool
}

func makeRequest(operation string, fo *Foaas) (string, error) {
	client := http.Client{}
	req, err := http.NewRequest(http.MethodGet, endpoint+operation, nil)
	if err != nil {
		return "", nil
	}

	var hval string
	switch fo.format {
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
	if fo.lang != "" {
		qry.Add("i18n", fo.lang)
	}
	if fo.shout == true {
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

func (fo *Foaas) Operations() (string, error) {
	return makeRequest("/operations", fo)
}

func (fo *Foaas) Anyway(company string, from string) (string, error) {
	url := fmt.Sprintf("/anyway/%s/%s", company, from)
	return makeRequest(url, fo)
}

func (fo *Foaas) Asshole(from string) (string, error) {
	url := fmt.Sprintf("/asshole/%s", from)
	return makeRequest(url, fo)
}

func (fo *Foaas) Awesome(from string) (string, error) {
	url := fmt.Sprintf("/awesome/%s", from)
	return makeRequest(url, fo)
}

func (fo *Foaas) Back(name string, from string) (string, error) {
	url := fmt.Sprintf("/back/%s/%s", name, from)
	return makeRequest(url, fo)
}

func (fo *Foaas) Bag(from string) (string, error) {
	url := fmt.Sprintf("/bag/%s", from)
	return makeRequest(url, fo)
}

func (fo *Foaas) Ballmer(name string, company string, from string) (string, error) {
	url := fmt.Sprintf("/ballmer/%s/%s/%s", name, company, from)
	return makeRequest(url, fo)
}

func (fo *Foaas) Birthday(name string, from string) (string, error) {
	url := fmt.Sprintf("/bday/%s/%s", name, from)
	return makeRequest(url, fo)
}

func (fo *Foaas) Because(from string) (string, error) {
	url := fmt.Sprintf("/because/%s", from)
	return makeRequest(url, fo)
}

func (fo *Foaas) Blackadder(name string, from string) (string, error) {
	url := fmt.Sprintf("/blackadder/%s/%s", name, from)
	return makeRequest(url, fo)
}

func (fo *Foaas) BravoMike(name string, from string) (string, error) {
	url := fmt.Sprintf("/bm/%s/%s", name, from)
	return makeRequest(url, fo)
}

func (fo *Foaas) Bucket(from string) (string, error) {
	url := fmt.Sprintf("/bucket/%s", from)
	return makeRequest(url, fo)
}

func (fo *Foaas) Bus(name string, from string) (string, error) {
	url := fmt.Sprintf("/bus/%s/%s", name, from)
	return makeRequest(url, fo)
}

func (fo *Foaas) Bye(from string) (string, error) {
	url := fmt.Sprintf("/bye/%s", from)
	return makeRequest(url, fo)
}

func (fo *Foaas) CanIUse(tool string, from string) (string, error) {
	url := fmt.Sprintf("/caniuse/%s/%s", tool, from)
	return makeRequest(url, fo)
}

func (fo *Foaas) Chainsaw(name string, from string) (string, error) {
	url := fmt.Sprintf("/chainsaw/%s/%s", name, from)
	return makeRequest(url, fo)
}

func (fo *Foaas) Cocksplat(name string, from string) (string, error) {
	url := fmt.Sprintf("/cocksplat/%s/%s", name, from)
	return makeRequest(url, fo)
}

func (fo *Foaas) Cool(from string) (string, error) {
	url := fmt.Sprintf("/cool/%s", from)
	return makeRequest(url, fo)
}

func (fo *Foaas) Cup(from string) (string, error) {
	url := fmt.Sprintf("/cup/%s", from)
	return makeRequest(url, fo)
}

func (fo *Foaas) Dalton(name string, from string) (string, error) {
	url := fmt.Sprintf("/dalton/%s/%s", name, from)
	return makeRequest(url, fo)
}

func (fo *Foaas) Deraadt(name string, from string) (string, error) {
	url := fmt.Sprintf("/deraadt/%s/%s", name, from)
	return makeRequest(url, fo)
}

func (fo *Foaas) Diabetes(from string) (string, error) {
	url := fmt.Sprintf("/diabetes/%s", from)
	return makeRequest(url, fo)
}

func (fo *Foaas) Donut(name string, from string) (string, error) {
	url := fmt.Sprintf("/donut/%s/%s", name, from)
	return makeRequest(url, fo)
}

func (fo *Foaas) DoSomething(do string, something string, from string) (string, error) {
	url := fmt.Sprintf("/dosomething/%s/%s/%s", do, something, from)
	return makeRequest(url, fo)
}

func (fo *Foaas) Equity(name string, from string) (string, error) {
	url := fmt.Sprintf("/equity/%s/%s", name, from)
	return makeRequest(url, fo)
}

func (fo *Foaas) Even(from string) (string, error) {
	url := fmt.Sprintf("/even/%s", from)
	return makeRequest(url, fo)
}

func (fo *Foaas) Everyone(from string) (string, error) {
	url := fmt.Sprintf("/everyone/%s", from)
	return makeRequest(url, fo)
}

func (fo *Foaas) Everything(from string) (string, error) {
	url := fmt.Sprintf("/everything/%s", from)
	return makeRequest(url, fo)
}

func (fo *Foaas) Family(from string) (string, error) {
	url := fmt.Sprintf("/family/%s", from)
	return makeRequest(url, fo)
}

func (fo *Foaas) Fascinating(from string) (string, error) {
	url := fmt.Sprintf("/fascinating/%s", from)
	return makeRequest(url, fo)
}

func (fo *Foaas) Fewer(name string, from string) (string, error) {
	url := fmt.Sprintf("/fewer/%s/%s", name, from)
	return makeRequest(url, fo)
}

func (fo *Foaas) Field(name string, from string, reference string) (string, error) {
	url := fmt.Sprintf("/field/%s/%s/%s", name, from, reference)
	return makeRequest(url, fo)
}

func (fo *Foaas) Flying(from string) (string, error) {
	url := fmt.Sprintf("/flying/%s", from)
	return makeRequest(url, fo)
}

func (fo *Foaas) Ftfy(from string) (string, error) {
	url := fmt.Sprintf("/ftfy/%s", from)
	return makeRequest(url, fo)
}

func (fo *Foaas) Fts(name string, from string) (string, error) {
	url := fmt.Sprintf("/fts/%s/%s", name, from)
	return makeRequest(url, fo)
}

func (fo *Foaas) Fyyff(from string) (string, error) {
	url := fmt.Sprintf("/fyyff/%s", from)
	return makeRequest(url, fo)
}

func (fo *Foaas) Gfy(name string, from string) (string, error) {
	url := fmt.Sprintf("/gfy/%s/%s", name, from)
	return makeRequest(url, fo)
}

func (fo *Foaas) Give(from string) (string, error) {
	url := fmt.Sprintf("/give/%s", from)
	return makeRequest(url, fo)
}

func (fo *Foaas) Greed(noun string, from string) (string, error) {
	url := fmt.Sprintf("/greed/%s/%s", noun, from)
	return makeRequest(url, fo)
}

func (fo *Foaas) HolyGrail(from string) (string, error) {
	url := fmt.Sprintf("/holygrail/%s", from)
	return makeRequest(url, fo)
}

func (fo *Foaas) Horse(from string) (string, error) {
	url := fmt.Sprintf("/horse/%s", from)
	return makeRequest(url, fo)
}

func (fo *Foaas) Idea(from string) (string, error) {
	url := fmt.Sprintf("/idea/%s", from)
	return makeRequest(url, fo)
}

func (fo *Foaas) Immensity(from string) (string, error) {
	url := fmt.Sprintf("/immensity/%s", from)
	return makeRequest(url, fo)
}

func (fo *Foaas) Ing(name string, from string) (string, error) {
	url := fmt.Sprintf("/ing/%s/%s", name, from)
	return makeRequest(url, fo)
}

func (fo *Foaas) JingleBells(from string) (string, error) {
	url := fmt.Sprintf("/jinglebells/%s", from)
	return makeRequest(url, fo)
}

func (fo *Foaas) Keep(name string, from string) (string, error) {
	url := fmt.Sprintf("/keep/%s/%s", name, from)
	return makeRequest(url, fo)
}

func (fo *Foaas) KeepCalm(reaction string, from string) (string, error) {
	url := fmt.Sprintf("/keepcalm/%s/%s", reaction, from)
	return makeRequest(url, fo)
}

func (fo *Foaas) King(name string, from string) (string, error) {
	url := fmt.Sprintf("/king/%s/%s", name, from)
	return makeRequest(url, fo)
}

func (fo *Foaas) Legend(name string, from string) (string, error) {
	url := fmt.Sprintf("/legend/%s/%s", name, from)
	return makeRequest(url, fo)
}

func (fo *Foaas) Life(from string) (string, error) {
	url := fmt.Sprintf("/life/%s", from)
	return makeRequest(url, fo)
}

func (fo *Foaas) Linus(from string) (string, error) {
	url := fmt.Sprintf("/linus/%s", from)
	return makeRequest(url, fo)
}

func (fo *Foaas) Logs(from string) (string, error) {
	url := fmt.Sprintf("/logs/%s", from)
	return makeRequest(url, fo)
}

func (fo *Foaas) Look(name string, from string) (string, error) {
	url := fmt.Sprintf("/look/%s/%s", name, from)
	return makeRequest(url, fo)
}

func (fo *Foaas) Looking(from string) (string, error) {
	url := fmt.Sprintf("/looking/%s", from)
	return makeRequest(url, fo)
}

func (fo *Foaas) Madison(name string, from string) (string, error) {
	url := fmt.Sprintf("/madison/%s/%s", name, from)
	return makeRequest(url, fo)
}

func (fo *Foaas) Maybe(from string) (string, error) {
	url := fmt.Sprintf("/Maybe/%s", from)
	return makeRequest(url, fo)
}

func (fo *Foaas) Morning(from string) (string, error) {
	url := fmt.Sprintf("/mornin/%s", from)
	return makeRequest(url, fo)
}

func (fo *Foaas) No(from string) (string, error) {
	url := fmt.Sprintf("/no/%s", from)
	return makeRequest(url, fo)
}

func (fo *Foaas) Nugget(name string, from string) (string, error) {
	url := fmt.Sprintf("/nugget/%s%s", name, from)
	return makeRequest(url, fo)
}

func (fo *Foaas) Off(name string, from string) (string, error) {
	url := fmt.Sprintf("/off/%s/%s", name, from)
	return makeRequest(url, fo)
}

func (fo *Foaas) OffWith(behavior string, from string) (string, error) {
	url := fmt.Sprintf("/off-with/%s/%s", behavior, from)
	return makeRequest(url, fo)
}

func (fo *Foaas) Outside(name string, from string) (string, error) {
	url := fmt.Sprintf("/outside/%s/%s", name, from)
	return makeRequest(url, fo)
}

func (fo *Foaas) Particular(thing string, from string) (string, error) {
	url := fmt.Sprintf("/particular/%s/%s", thing, from)
	return makeRequest(url, fo)
}

func (fo *Foaas) Ping(from string) (string, error) {
	url := fmt.Sprintf("/pink/%s", from)
	return makeRequest(url, fo)
}

func (fo *Foaas) Problem(name string, from string) (string, error) {
	url := fmt.Sprintf("/problem/%s/%s", name, from)
	return makeRequest(url, fo)
}

func (fo *Foaas) Programmer(from string) (string, error) {
	url := fmt.Sprintf("/programmer/%s", from)
	return makeRequest(url, fo)
}

func (fo *Foaas) Pulp(language string, from string) (string, error) {
	url := fmt.Sprintf("/pulp/%s/%s", language, from)
	return makeRequest(url, fo)
}

func (fo *Foaas) Question(from string) (string, error) {
	url := fmt.Sprintf("/question/%s", from)
	return makeRequest(url, fo)
}

func (fo *Foaas) RatsArse(from string) (string, error) {
	url := fmt.Sprintf("/typehere/%s", from)
	return makeRequest(url, fo)
}

func (fo *Foaas) Retard(from string) (string, error) {
	url := fmt.Sprintf("/retard/%s", from)
	return makeRequest(url, fo)
}

func (fo *Foaas) Ridiculous(from string) (string, error) {
	url := fmt.Sprintf("/ridiculous/%s", from)
	return makeRequest(url, fo)
}

func (fo *Foaas) Rockstar(name string, from string) (string, error) {
	url := fmt.Sprintf("/rockstar/%s/%s", name, from)
	return makeRequest(url, fo)
}

func (fo *Foaas) Rtfm(from string) (string, error) {
	url := fmt.Sprintf("/rtfm/%s", from)
	return makeRequest(url, fo)
}

func (fo *Foaas) Sake(from string) (string, error) {
	url := fmt.Sprintf("/sake/%s", from)
	return makeRequest(url, fo)
}

func (fo *Foaas) Shakespeare(name string, from string) (string, error) {
	url := fmt.Sprintf("/shakespeare/%s/%s", name, from)
	return makeRequest(url, fo)
}

func (fo *Foaas) Shit(from string) (string, error) {
	url := fmt.Sprintf("/shit/%s", from)
	return makeRequest(url, fo)
}

func (fo *Foaas) ShutUp(name string, from string) (string, error) {
	url := fmt.Sprintf("/shutup/%s/%s", name, from)
	return makeRequest(url, fo)
}

func (fo *Foaas) Single(from string) (string, error) {
	url := fmt.Sprintf("/single/%s", from)
	return makeRequest(url, fo)
}

func (fo *Foaas) Thanks(from string) (string, error) {
	url := fmt.Sprintf("/thanks/%s", from)
	return makeRequest(url, fo)
}

func (fo *Foaas) That(from string) (string, error) {
	url := fmt.Sprintf("/that/%s", from)
	return makeRequest(url, fo)
}

func (fo *Foaas) Think(name string, from string) (string, error) {
	url := fmt.Sprintf("/think/%s/%s", name, from)
	return makeRequest(url, fo)
}

func (fo *Foaas) Thinking(name string, from string) (string, error) {
	url := fmt.Sprintf("/thinking/%s/%s", name, from)
	return makeRequest(url, fo)
}

func (fo *Foaas) This(from string) (string, error) {
	url := fmt.Sprintf("/this/%s", from)
	return makeRequest(url, fo)
}

func (fo *Foaas) Thumbs(name string, from string) (string, error) {
	url := fmt.Sprintf("/thumbs/%s/%s", name, from)
	return makeRequest(url, fo)
}

func (fo *Foaas) Too(from string) (string, error) {
	url := fmt.Sprintf("/too/%s", from)
	return makeRequest(url, fo)
}

func (fo *Foaas) Tucker(from string) (string, error) {
	url := fmt.Sprintf("/tucker/%s", from)
	return makeRequest(url, fo)
}

func (fo *Foaas) Version() (string, error) {
	return makeRequest("/version", fo)
}

func (fo *Foaas) Waste(name string, from string) (string, error) {
	url := fmt.Sprintf("/waste/%s/%s", name, from)
	return makeRequest(url, fo)
}

func (fo *Foaas) What(from string) (string, error) {
	url := fmt.Sprintf("/what/%s", from)
	return makeRequest(url, fo)
}

func (fo *Foaas) Xmas(name string, from string) (string, error) {
	url := fmt.Sprintf("/xmas/%s/%s", name, from)
	return makeRequest(url, fo)
}

func (fo *Foaas) Yoda(name string, from string) (string, error) {
	url := fmt.Sprintf("/yoda/%s/%s", name, from)
	return makeRequest(url, fo)
}

func (fo *Foaas) You(name string, from string) (string, error) {
	url := fmt.Sprintf("/you/%s/%s", name, from)
	return makeRequest(url, fo)
}

func (fo *Foaas) Zayn(from string) (string, error) {
	url := fmt.Sprintf("/zayn/%s", from)
	return makeRequest(url, fo)
}

func (fo *Foaas) Zero(from string) (string, error) {
	url := fmt.Sprintf("/zero/%s", from)
	return makeRequest(url, fo)
}
