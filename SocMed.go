package main

import "fmt"

type user struct {
	ID       int
	username string
	password string
	status   string
	friend   [20]person
	nFriend  int
}

type person struct {
	username   string
	status     string
	comments   [20]string
	statFriend bool
}

type tabPerson [20]person

type tabUser [100]user

type Application struct {
	Users    []user
	LoggedIn user
	NextID   int
	StatusID int
	FriendID int
}

func Run() {
	app := Application{} // Create an instance of the Application struct

	for {
		fmt.Println("==== Social Media App ====")
		fmt.Println("1. Register")
		fmt.Println("2. Login")
		fmt.Println("3. Exit")
		var menuOption int
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&menuOption)

		switch menuOption {
		case 1:
			register(&app) // Pass the app instance to the register function
		case 2:
			login(&app) // Pass the app instance to the login function
		case 3:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func main() {
	Run()
}

func login(app *Application) {
	fmt.Println("==== Login ====")
	var username, password string

	fmt.Print("Enter username: ")
	fmt.Scanln(&username)

	fmt.Print("Enter password: ")
	fmt.Scanln(&password)

	for _, user := range app.Users {
		if user.username == username && user.password == password {
			app.LoggedIn = user
			homeMenu(app) // Call the homeMenu function with the app instance
			return
		}
	}

	fmt.Println("Invalid username or password.")
}

func register(app *Application) {
	fmt.Println("==== Register ====")
	var username, password string

	fmt.Print("Enter username: ")
	fmt.Scanln(&username)

	fmt.Print("Enter password: ")
	fmt.Scanln(&password)

	user := user{
		ID:       app.NextID,
		username: username,
		password: password,
	}

	app.NextID++
	app.Users = append(app.Users, user)
	app.Users[app.NextID-1].ID = app.NextID - 1

	fmt.Println("Registration successful.")
}

func homeMenu(app *Application) {
	var orang tabPerson
	// var ID int
	// var userUsername string
	isiPerson(&orang)

	// ID = searchUsers(&app, userUsername)
	for {
		fmt.Println("==== Home Menu ====")
		fmt.Println("1. View All Statuses")
		fmt.Println("2. Add Comment")
		fmt.Println("3. Add Friend")
		fmt.Println("4. Remove Friend")
		fmt.Println("5. Edit Profile")
		fmt.Println("6. View Friends")
		fmt.Println("7. Search Users")
		fmt.Println("8. Logout")
		var menuOption int
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&menuOption)

		switch menuOption {
		case 1:
			viewAllStatus(&app.LoggedIn, orang)
		case 2:
			addComment(&app.Users[app.LoggedIn.ID], orang)
		case 3:
			addFriend(&app.Users[app.LoggedIn.ID], orang)
		case 4:
			deleteFriend(&app.Users[app.LoggedIn.ID], orang)
		case 5:
			editProfile(&app.Users[app.LoggedIn.ID])
		case 6:
			urut(&app.Users[app.LoggedIn.ID])
			ViewFriends(app.Users[app.LoggedIn.ID])
		case 7:
			searchPersonReal(orang)
		case 8:
			fmt.Println("Logging out...")
			//!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
			fmt.Println(app.Users[app.LoggedIn.ID].username)
			fmt.Println(app.Users[app.LoggedIn.ID].password)
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
func viewAllStatus(pengguna *user, orang tabPerson) {
	fmt.Println("==== View All Statuses ====")
	if len(orang) == 0 {
		fmt.Println("No statuses found.")
		return
	}

	for i, person := range orang {
		fmt.Printf("[%d]\n", i+1)
		fmt.Printf("Username: %s\n", person.username)
		fmt.Printf("Status: %s\n", person.status)
		fmt.Println("===================================")
	}
}

func addComment(pengguna *user, orang tabPerson) {
	fmt.Println("==== Add Comment ====")
	fmt.Print("Enter your comment: ")
	var comment string
	fmt.Scanln(&comment)
	fmt.Println("Comment added successfully.")
}

func addFriend(pengguna *user, orang tabPerson) {
	var friendUsername string
	var found int = -1

	fmt.Println("==== Add Friend ====")
	fmt.Print("Enter username: ")
	fmt.Scanln(&friendUsername)

	found = searchPerson(orang, friendUsername)
	for found == -1 {
		fmt.Println("Invalid username")
		fmt.Print("Enter username: ")
		fmt.Scanln(&friendUsername)
		found = searchPerson(orang, friendUsername)
	}

	if searchFriend(*pengguna, friendUsername) == -1 {
		pengguna.friend[pengguna.nFriend].username = orang[found].username
		pengguna.nFriend += 1

		fmt.Println("Friend added successfully.")
	} else {
		fmt.Println("You already add", friendUsername)
	}
}

func deleteFriend(pengguna *user, orang tabPerson) {
	var friendUsername string
	var found int = -1

	fmt.Println("==== Delete Friend ====")
	fmt.Print("Enter username: ")
	fmt.Scanln(&friendUsername)

	//buat search frienddd
	found = searchFriend(*pengguna, friendUsername)
	for found == -1 {
		fmt.Println("Invalid username")
		fmt.Print("Enter username: ")
		fmt.Scanln(&friendUsername)
		found = searchFriend(*pengguna, friendUsername)
	}
	//buat ngegeser
	delete(found, &*pengguna)

	fmt.Println("Friend removed successfully.")
}

func delete(found int, pengguna *user) {
	var n, i int

	i = found
	n = pengguna.nFriend
	for i <= n-1 {
		pengguna.friend[i] = pengguna.friend[i+1]
		i++
	}
	pengguna.nFriend -= 1
}

// function u/ mencari person. dipakai di add friend
func searchPerson(orang tabPerson, friendUsername string) int {
	var hasil int = -1
	var i int = 0

	for hasil == -1 && i < 20 {
		if orang[i].username == friendUsername {
			hasil = i
		}
		i++
	}
	return hasil
}

func searchPersonReal(orang tabPerson) {
	var friendUsername string
	var found int = -1

	fmt.Println("==== Search Person ====")
	fmt.Print("Enter username: ")
	fmt.Scanln(&friendUsername)

	//buat search frienddd
	found = searchPerson(*&orang, friendUsername)
	for found == -1 {
		fmt.Println("Username not found.")
		fmt.Print("Enter username: ")
		fmt.Scanln(&friendUsername)
		found = searchPerson(*&orang, friendUsername)
	}

	fmt.Println("Username found.")
}

// function u/ mencari teman. dipakai di delete friend
func searchFriend(pengguna user, friendUsername string) int {
	var hasil int = -1
	var i int = 0

	for hasil == -1 && i < 20 {
		if pengguna.friend[i].username == friendUsername {
			hasil = i
		}
		i++
	}
	return hasil
}

func searchUsers(app *Application, userUsername string) int {
	var hasil int = -1
	var i int = 0

	for hasil == -1 && i < 20 {
		if app.Users[i].username == userUsername {
			hasil = i
		}
		i++
	}
	return hasil
}

// masukin nama" orang nya DUMMY -----------
func isiPerson(person *tabPerson) {
	person[0].username = "AlexAdventures"
	person[0].status = "Halo, apa kabar?"
	person[1].username = "ZahraAmiera"
	person[1].status = "Gua laper"
	person[2].username = "MaxMusings"
	person[2].status = "Selamat ulang tahun!"
	person[3].username = "LilyLovesLife"
	person[3].status = "Selamat pagi!"
	person[4].username = "EthanEnthusiast"
	person[4].status = "Sampai jumpa!"
	person[5].username = "MarchianoAulia"
	person[5].status = "Bagaimana keadaannya?"
	person[6].username = "NoahNarratives"
	person[6].status = "Tolong bantu saya."
	person[7].username = "AvaAdventurer"
	person[7].status = "Maju Maju"
	person[8].username = "LucasLifeJourney"
	person[8].status = "Apa Kabar ?"
	person[9].username = "SophiaStoryteller"
	person[9].status = "Hello World !"
	person[10].username = "BenjaminBliss"
	person[10].status = "Hi there :)"
	person[11].username = "MiaMoments"
	person[11].status = "Good Morning :) "
	person[12].username = "SiapaNggaTau"
	person[12].status = "loh gk tau kok tanya saya"
	person[13].username = "GraceGlobeTrotter"
	person[13].status = "Wow! Apakah itu bukan "
	person[14].username = "JamesJourneys"
	person[14].status = "Kapan lagi aku dateng? "
	person[15].username = "IsabellaImagines"
	person[15].status = "ngadi ngadi"
	person[16].username = "OliverOdyssey"
	person[16].status = "I'm here to help you with your problem"
	person[17].username = "AmeliaAdventures"
	person[17].status = "What's up guys?? "
	person[18].username = "JackJourno"
	person[18].status = "How are u doing today ??"
	person[19].username = "HarperHighlights"
	person[19].status = "Haloo..!! Selamat Pagi.."

	//isi status
	for i := 0; i < 20; i++ {
		person[i].statFriend = false
	}
}

func ViewFriends(teman user) {
	var masihAda bool = true

	fmt.Println("==== Friends ====")
	for i := 0; i < 20 && masihAda; i++ {
		fmt.Println(teman.friend[i].username)
		masihAda = teman.friend[i+1].username != ""
	}
	fmt.Println()
}

func urut(set *user) {
	//insertion sort
	var pass, i, n int
	var temp string

	n = set.nFriend
	pass = 1
	for pass <= n-1 {
		i = pass
		temp = set.friend[pass].username
		for i > 0 && temp < set.friend[i-1].username {
			set.friend[i].username = set.friend[i-1].username
			i = i - 1
		}
		set.friend[i].username = temp
		pass += 1
	}
}

func editProfile(pengguna *user) {
	var inputPilihan int
	fmt.Println("==== Edit Profile ====")
	menuEditProfile()
	fmt.Scan(&inputPilihan)
	for !(inputPilihan == 1 || inputPilihan == 2) {
		fmt.Println("Invalid choice. Please try again")
		menuEditProfile()
		fmt.Scan(&inputPilihan)
	}
	if inputPilihan == 1 {
		changeUserUsername(&*pengguna)
	} else if inputPilihan == 2 {
		changeUserPassword(&*pengguna)
	}
}

func menuEditProfile() {
	fmt.Println("1. Edit Username")
	fmt.Println("2. Edit Password")
	fmt.Print("Enter your choice: ")
}

func changeUserUsername(pengguna *user) {
	var newUsername string

	fmt.Print("Enter new username: ")
	fmt.Scan(&newUsername)
	pengguna.username = newUsername

	fmt.Println("Username changed successfully")
	fmt.Println(pengguna.username)
}

func changeUserPassword(pengguna *user) {
	var newPassword string

	fmt.Print("Enter new password: ")
	fmt.Scan(&newPassword)
	pengguna.password = newPassword

	fmt.Println("Password changed successfully")
}
