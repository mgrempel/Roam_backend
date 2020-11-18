package utilities

import "Roam/Roam_backend/graph/model"

//ScrubUser workaround for stopping the API from coughing up sensitive stuff.
// I'll implement a better solution once the project is more fleshed out
func ScrubUser(user *model.User) {
	user.Password = "Nice try."
	user.UUID = "Not gonna happen."
}
