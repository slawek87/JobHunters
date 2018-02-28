package user

import "testing"

func TestUserGetters(t *testing.T) {
	user := &User{UserID: "xyz123", FirstName: "Tom", LastName: "Gardner", Avatar: "photo.png", Email: "tom@tom.tom"}

	if GetID(user) != user.UserID {
		t.Error("GetFirstName doesn't return correct value.")
	}

	if GetFirstName(user) != user.FirstName {
		t.Error("GetFirstName doesn't return correct value.")
	}

	if GetLastName(user) != user.LastName {
		t.Error("GetLastName doesn't return correct value.")
	}

	if GetAvatar(user) != user.Avatar {
		t.Error("Avatar doesn't return correct value.")
	}

	if GetEmail(user) != user.Email {
		t.Error("Email doesn't return correct value.")
	}
}
