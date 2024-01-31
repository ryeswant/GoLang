package main

import "fmt"

type OrganizingTeam struct {
	TeamMembers    []string
	PrimaryContact string
}

type NSPEvent struct {
	EventDate      string
	EventName      string
	OrganizingTeam OrganizingTeam
}

func NewNSPEvent(date, name string, teamMembers []string, primaryContact string) *NSPEvent {
	event := &NSPEvent{
		EventDate: date,
		EventName: name,
		OrganizingTeam: OrganizingTeam{
			TeamMembers:    teamMembers,
			PrimaryContact: primaryContact,
		},
	}
	return event
}

func main() {
	event := NewNSPEvent("2024-01-15", "NSP Sankranthi Event", []string{"Yeswanth", "Rakesh", "Harshitha"}, "John Doe")

	event.OrganizingTeam.TeamMembers = append(event.OrganizingTeam.TeamMembers, "Deepak")
	event.OrganizingTeam.PrimaryContact = "Rakesh"

	fmt.Println("Event Details:")
	fmt.Println("Date:", event.EventDate)
	fmt.Println("Name:", event.EventName)
	fmt.Println("Organizing Team Members:", event.OrganizingTeam.TeamMembers)
	fmt.Println("Organizing Team Primary Contact:", event.OrganizingTeam.PrimaryContact)
}
