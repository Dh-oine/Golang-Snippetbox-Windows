package main

import (
	"fmt"
	"net/http"
	"snippetbox/pkg/forms"
	"snippetbox/pkg/models"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	//if r.URL.Path != "/" {
	//	app.notFound(w)
	//	return
	//}

	s, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	//data := &templateData{Snippets: s}
	//
	//files := []string{
	//	"C:/Users/Богдан/code/snippetbox/ui/html/home.page.tmpl",
	//	"C:/Users/Богдан/code/snippetbox/ui/html/base.layout.tmpl",
	//	"C:/Users/Богдан/code/snippetbox/ui/html/footer.partial.tmpl",
	//}
	//
	//ts, err := template.ParseFiles(files...)
	//if err != nil {
	//	app.serverError(w, err)
	//	return
	//}
	//err = ts.Execute(w, data)
	//if err != nil {
	//	app.serverError(w, err)
	//}
	app.render(w, r, "home.page.tmpl", &templateData{
		Snippets: s,
	})
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	s, err := app.snippets.Get(id)
	if err == models.ErrNoRecord {
		app.notFound(w)
		return
	} else if err != nil {
		app.serverError(w, err)
		return
	}
	//flash := app.session.PopString(r, "flash")

	//data := &templateData{Snippet: s}
	//	//
	//	//files := []string{
	//	//	"C:/Users/Богдан/code/snippetbox/ui/html/show.page.tmpl",
	//	//	"C:/Users/Богдан/code/snippetbox/ui/html/base.layout.tmpl",
	//	//	"C:/Users/Богдан/code/snippetbox/ui/html/footer.partial.tmpl",
	//	//}
	//	//
	//	//ts, err := template.ParseFiles(files...)
	//	//
	//	//if err != nil {
	//	//	app.serverError(w, err)
	//	//	return
	//	//}
	//	//
	//	//err = ts.Execute(w, data)
	//	//if err != nil {
	//	//	app.serverError(w, err)
	//	//}

	app.render(w, r, "show.page.tmpl", &templateData{
		//Flash:   flash,
		Snippet: s,
	})
}

func (app *application) createSnippetForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "create.page.tmpl", &templateData{
		Form: forms.New(nil),
	})
	//app.render(w, r, "create.page.tmpl", nil)
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	form := forms.New(r.PostForm)
	form.Required("title", "content", "expires")
	form.MaxLength("title", 100)
	form.PermittedValues("expires", "365", "7", "1")
	if !form.Valid() {
		app.render(w, r, "create.page.tmpl", &templateData{Form: form})
		return
	}
	id, err := app.snippets.Insert(form.Get("title"), form.Get("content"), form.Get("expires"))
	if err != nil {
		app.serverError(w, err)
		return
	}
	app.session.Put(r, "flash", "Snippet successfully created!")
	//err := r.ParseForm()
	//if err != nil {
	//	app.clientError(w, http.StatusBadRequest)
	//	return
	//}
	//
	////title := r.PostForm.Get("title")
	////content := r.PostForm.Get("content")
	////expires := r.PostForm.Get("expires")
	//form := forms.New(r.PostForm)
	//form.Required("title", "content", "expires")
	//form.MaxLength("title", 100)
	//form.PermittedValues("expires", "365", "7", "1")
	//if !form.Valid() {
	//	app.render(w, r, "create.page.tmpl", &templateData{Form: form})
	//}
	//id, err := app.snippets.Insert(form.Get("title"), form.Get("content"), form.Get("expires"))
	//if err != nil {
	//	app.serverError(w, err)
	//	return
	//}
	//errors := make(map[string]string)
	//
	//if strings.TrimSpace(title) == "" {
	//	errors["title"] = "This field cannot be blank"
	//} else if utf8.RuneCountInString(title) > 100 {
	//	errors["title"] = "This field is too long (maximum is 100 characters)"
	//}
	//
	//if strings.TrimSpace(content) == "" {
	//	errors["content"] = "This field cannot be blank"
	//}
	//
	//if strings.TrimSpace(expires) == "" {
	//	errors["expires"] = "This field cannot be blank"
	//} else if expires != "365" && expires != "7" && expires != "1" {
	//	errors["expires"] = "This field is invalid"
	//}
	//if len(errors) > 0 {
	//	app.render(w, r, "create.page.tmpl", &templateData{
	//		FormErrors: errors,
	//		FormData:   r.PostForm,
	//	})
	//	return
	//}
	//
	//id, err := app.snippets.Insert(title, content, expires)
	//if err != nil {
	//	app.serverError(w, err)
	//	return
	//}
	//
	http.Redirect(w, r, fmt.Sprintf("/snippet/%d", id), http.StatusSeeOther)
}
func (app *application) signupUserForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "signup.page.tmpl", &templateData{
		Form: forms.New(nil),
	})
}
func (app *application) signupUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	form := forms.New(r.PostForm)
	form.Required("name", "email", "password")
	form.MatchesPattern("email", forms.EmailRX)
	form.MinLength("password", 10)
	if !form.Valid() {
		app.render(w, r, "signup.page.tmpl", &templateData{Form: form})
		return
	}
	err = app.users.Insert(form.Get("name"), form.Get("email"), form.Get("password"))
	if err == models.ErrDuplicateEmail {
		form.Errors.Add("email", "Address is already in use")
		app.render(w, r, "signup.page.tmpl", &templateData{Form: form})
		return
	} else if err != nil {
		app.serverError(w, err)
		return
	}
	app.session.Put(r, "flash", "Your signup was successful. Please log in.")

	http.Redirect(w, r, "/user/login", http.StatusSeeOther)
}
func (app *application) loginUserForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "login.page.tmpl", &templateData{
		Form: forms.New(nil),
	})
}
func (app *application) loginUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	form := forms.New(r.PostForm)
	id, err := app.users.Authenticate(form.Get("email"), form.Get("password"))
	if err == models.ErrInvalidCredentials {
		form.Errors.Add("generic", "Email or Password is incorrect")
		app.render(w, r, "login.page.tmpl", &templateData{Form: form})
		return
	} else if err != nil {
		app.serverError(w, err)
		return
	}
	app.session.Put(r, "userID", id)
	http.Redirect(w, r, "/snippet/create", http.StatusSeeOther)
}
func (app *application) logoutUser(w http.ResponseWriter, r *http.Request) {
	app.session.Remove(r, "userID")
	app.session.Put(r, "flash", "You've benn logged out successfully!")
	http.Redirect(w, r, "/", 303)
}
func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
