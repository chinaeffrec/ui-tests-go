package pages

import (
	"github.com/tebeka/selenium"
)

type FormPage struct {
	Driver selenium.WebDriver
	URI    string
}

const URL = "https://practice-automation.com/form-fields/"

func NewFormPage(driver selenium.WebDriver) *FormPage {
	return &FormPage{
		Driver: driver,
		URI:    URL,
	}
}

func (p *FormPage) OpenPage() error {
	if err := p.Driver.Get(p.URI); err != nil {
		return err
	}
	return nil
}

func (p *FormPage) EnterName(name string) error {
	element, err := p.Driver.FindElement(selenium.ByID, "name-input")
	if err != nil {
		return err
	}
	element.Clear()
	return element.SendKeys(name)
}

func (p *FormPage) EnterPassword(password string) error {
	element, err := p.Driver.FindElement(selenium.ByCSSSelector, "input[type='password']")
	if err != nil {
		return err
	}
	element.Clear()
	return element.SendKeys(password)
}

func (p *FormPage) SelectDrinks(drinkIDs ...string) error {
	for _, id := range drinkIDs {
		element, err := p.Driver.FindElement(selenium.ByID, id)
		if err != nil {
			return err
		}
		if err = element.Click(); err != nil {
			return err
		}
	}
	return nil
}

func (p *FormPage) SelectColor(colorID string) error {
	element, err := p.Driver.FindElement(selenium.ByID, colorID)
	if err != nil {
		return err
	}
	return element.Click()
}

func (p *FormPage) SelectAutomation(optionValue string) error {
	element, err := p.Driver.FindElement(selenium.ByID, "automation")
	if err != nil {
		return err
	}
	return element.SendKeys(optionValue)
}

func (p *FormPage) EnterEmail(email string) error {
	element, err := p.Driver.FindElement(selenium.ByID, "email")
	if err != nil {
		return err
	}
	return element.SendKeys(email)
}

func (p *FormPage) EnterMessage(message string) error {
	element, err := p.Driver.FindElement(selenium.ByID, "message")
	if err != nil {
		return err
	}
	return element.SendKeys(message)
}

func (p *FormPage) Submit() error {
	element, err := p.Driver.FindElement(selenium.ByID, "submit-btn")
	if err != nil {
		return err
	}
	return element.Click()
}
