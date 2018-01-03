package BaseController

import (
	"math/rand"
	"os"
	"time"

	"../../models/UserModel"
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
)

// Base is super struct
type Base struct {
	Ctx     iris.Context
	Session *sessions.Session
}

// LoginUser is login
func (c *Base) LoginUser(username string) error {
	userModel := UserModel.New()
	userid, teamname, teamid, err := userModel.GetUserInfo(username)
	if err != nil {
		return err
	}
	c.Session.Set("username", username)
	c.Session.Set("userid", userid)
	c.Session.Set("teamname", teamname)
	c.Session.Set("teamid", teamid)
	return nil
}

// IsLoggedIn is Check login status
func (c *Base) IsLoggedIn() bool {
	return c.Session.Get("username") != nil
}

// GetLoggedUserName is Check login status
func (c *Base) GetLoggedUserName() string {
	return c.Session.GetString("username")
}

// GetLoggedUserID is return user id
func (c *Base) GetLoggedUserID() string {
	return c.Session.GetString("userid")
}

// GetLoggedTeamName is return team id
func (c *Base) GetLoggedTeamName() string {
	return c.Session.GetString("teamname")
}

// GetLoggedTeamID is return team id
func (c *Base) GetLoggedTeamID() string {
	return c.Session.GetString("teamid")
}

// Logout is logout
func (c *Base) Logout() {
	c.Session.Delete("username")
	c.Session.Delete("userid")
	c.Session.Delete("teamname")
	c.Session.Delete("teamid")
}

// MakeToken is generate taken and set taken in session.
func (c *Base) MakeToken() string {
	var rs1Letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ$%&")
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, 64)
	for i := range b {
		b[i] = rs1Letters[rand.Intn(len(rs1Letters))]
	}
	taken := string(b)
	c.Session.Set("_csrfToken", taken)
	return taken
}

// CheckTaken is check taken
func (c *Base) CheckTaken(taken string) bool {
	r := c.Session.Get("_csrfToken") == taken
	c.Session.Delete("_csrfToken")
	return r
}

func (c *Base) IsNowCompetition() bool {
	nowTime := time.Now().Unix()
	startTime, _ := time.Parse("2006-01-02 15:04:05" /*layout*/, os.Getenv("COMPETITION_START_TIME"))
	endTime, _ := time.Parse("2006-01-02 15:04:05" /*layout*/, os.Getenv("COMPETITION_END_TIME"))
	if startTime.Unix()-32400 < nowTime && nowTime < endTime.Unix()-32400 {
		return true
	}
	return false
}

func (c *Base) IsBeforeCompetition() bool {
	nowTime := time.Now().Unix()
	startTime, _ := time.Parse("2006-01-02 15:04:05" /*layout*/, os.Getenv("COMPETITION_START_TIME"))
	if startTime.Unix()-32400 > nowTime {
		return true
	}
	return false
}
