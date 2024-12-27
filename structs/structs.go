package structs

type User struct {
    LinkedInURN string
    TwitterID   string 
    AccessToken string 
}

type PostMessage struct {
    User       User   
    Body       string 
    Visibility string 
    Platform   string 
}
