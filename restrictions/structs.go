package restrictions

// Apply Restriction
type UsersToRestrict struct {
	UserId, Rule string
}

type RestrictionApplied []struct {
	Message string `json:"message"`
	UserID  int    `json:"user_id"`
}


// Release
type UsersToRelease struct {
	UserId string
}

type ReleasementApplied struct {
	Applied       bool        `json:"applied"`
	Reason        string 	  `json:"reason"`
	NewUserStatus string      `json:"new_user_status"`
}

/*
{
    "applied": true,
    "reason": null,
    "new_user_status": "init"
}
*/