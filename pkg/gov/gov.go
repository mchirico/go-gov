package gov

import (
	"encoding/json"
	"fmt"
	"github.com/mchirico/go-gov/pkg/httputils"
	"sync"
)

type GOV struct {
	Status    string `json:"status"`
	Copyright string `json:"copyright"`
	Results   []struct {
		Congress   string `json:"congress"`
		Chamber    string `json:"chamber"`
		NumResults int    `json:"num_results"`
		Offset     int    `json:"offset"`
		Members    []struct {
			ID                   string      `json:"id"`
			Title                string      `json:"title"`
			ShortTitle           string      `json:"short_title"`
			APIURI               string      `json:"api_uri"`
			FirstName            string      `json:"first_name"`
			MiddleName           interface{} `json:"middle_name"`
			LastName             string      `json:"last_name"`
			Suffix               interface{} `json:"suffix"`
			DateOfBirth          string      `json:"date_of_birth"`
			Gender               string      `json:"gender"`
			Party                string      `json:"party"`
			LeadershipRole       interface{} `json:"leadership_role"`
			TwitterAccount       string      `json:"twitter_account"`
			FacebookAccount      string      `json:"facebook_account"`
			YoutubeAccount       string      `json:"youtube_account"`
			GovtrackID           string      `json:"govtrack_id"`
			CspanID              string      `json:"cspan_id"`
			VotesmartID          string      `json:"votesmart_id"`
			IcpsrID              string      `json:"icpsr_id"`
			CrpID                string      `json:"crp_id"`
			GoogleEntityID       string      `json:"google_entity_id"`
			FecCandidateID       string      `json:"fec_candidate_id"`
			URL                  string      `json:"url"`
			RssURL               string      `json:"rss_url"`
			ContactForm          string      `json:"contact_form"`
			InOffice             bool        `json:"in_office"`
			CookPvi              interface{} `json:"cook_pvi"`
			DwNominate           float64     `json:"dw_nominate"`
			IdealPoint           interface{} `json:"ideal_point"`
			Seniority            string      `json:"seniority"`
			NextElection         string      `json:"next_election"`
			TotalVotes           int         `json:"total_votes"`
			MissedVotes          int         `json:"missed_votes"`
			TotalPresent         int         `json:"total_present"`
			LastUpdated          string      `json:"last_updated"`
			OcdID                string      `json:"ocd_id"`
			Office               string      `json:"office"`
			Phone                string      `json:"phone"`
			Fax                  string      `json:"fax"`
			State                string      `json:"state"`
			SenateClass          string      `json:"senate_class"`
			StateRank            string      `json:"state_rank"`
			LisID                string      `json:"lis_id"`
			MissedVotesPct       float64     `json:"missed_votes_pct"`
			VotesWithPartyPct    float64     `json:"votes_with_party_pct"`
			VotesAgainstPartyPct float64     `json:"votes_against_party_pct"`
		} `json:"members"`
	} `json:"results"`
}

type Subcommittees struct {
	Status    string `json:"status"`
	Copyright string `json:"copyright"`
	Results   []struct {
		ID              string      `json:"id"`
		MemberID        string      `json:"member_id"`
		FirstName       string      `json:"first_name"`
		MiddleName      interface{} `json:"middle_name"`
		LastName        string      `json:"last_name"`
		Suffix          interface{} `json:"suffix"`
		DateOfBirth     string      `json:"date_of_birth"`
		Gender          string      `json:"gender"`
		URL             string      `json:"url"`
		TimesTopicsURL  string      `json:"times_topics_url"`
		TimesTag        string      `json:"times_tag"`
		GovtrackID      string      `json:"govtrack_id"`
		CspanID         string      `json:"cspan_id"`
		VotesmartID     string      `json:"votesmart_id"`
		IcpsrID         string      `json:"icpsr_id"`
		TwitterAccount  string      `json:"twitter_account"`
		FacebookAccount string      `json:"facebook_account"`
		YoutubeAccount  interface{} `json:"youtube_account"`
		CrpID           string      `json:"crp_id"`
		GoogleEntityID  string      `json:"google_entity_id"`
		RssURL          string      `json:"rss_url"`
		InOffice        bool        `json:"in_office"`
		CurrentParty    string      `json:"current_party"`
		MostRecentVote  string      `json:"most_recent_vote"`
		LastUpdated     string      `json:"last_updated"`
		Roles           []struct {
			Congress             string      `json:"congress"`
			Chamber              string      `json:"chamber"`
			Title                string      `json:"title"`
			ShortTitle           string      `json:"short_title"`
			State                string      `json:"state"`
			Party                string      `json:"party"`
			LeadershipRole       string      `json:"leadership_role"`
			FecCandidateID       string      `json:"fec_candidate_id"`
			Seniority            string      `json:"seniority"`
			District             string      `json:"district"`
			AtLarge              bool        `json:"at_large"`
			OcdID                string      `json:"ocd_id"`
			StartDate            string      `json:"start_date"`
			EndDate              string      `json:"end_date"`
			Office               string      `json:"office"`
			Phone                string      `json:"phone"`
			Fax                  interface{} `json:"fax"`
			ContactForm          interface{} `json:"contact_form"`
			CookPvi              string      `json:"cook_pvi"`
			DwNominate           float64     `json:"dw_nominate"`
			IdealPoint           interface{} `json:"ideal_point"`
			NextElection         string      `json:"next_election"`
			TotalVotes           int         `json:"total_votes"`
			MissedVotes          int         `json:"missed_votes"`
			TotalPresent         int         `json:"total_present"`
			SenateClass          string      `json:"senate_class"`
			StateRank            string      `json:"state_rank"`
			LisID                string      `json:"lis_id"`
			BillsSponsored       int         `json:"bills_sponsored"`
			BillsCosponsored     int         `json:"bills_cosponsored"`
			MissedVotesPct       float64     `json:"missed_votes_pct"`
			VotesWithPartyPct    float64     `json:"votes_with_party_pct"`
			VotesAgainstPartyPct float64     `json:"votes_against_party_pct"`
			Committees           []struct {
				Name        string `json:"name"`
				Code        string `json:"code"`
				APIURI      string `json:"api_uri"`
				Side        string `json:"side"`
				Title       string `json:"title"`
				RankInParty int    `json:"rank_in_party"`
				BeginDate   string `json:"begin_date"`
				EndDate     string `json:"end_date"`
			} `json:"committees"`
			Subcommittees []struct {
				Name              string `json:"name"`
				Code              string `json:"code"`
				ParentCommitteeID string `json:"parent_committee_id"`
				APIURI            string `json:"api_uri"`
				Side              string `json:"side"`
				Title             string `json:"title"`
				RankInParty       int    `json:"rank_in_party"`
				BeginDate         string `json:"begin_date"`
				EndDate           string `json:"end_date"`
			} `json:"subcommittees"`
		} `json:"roles"`
	} `json:"results"`
}

type Gov struct {
	Key map[string]string
	sync.Mutex
}

func NewGov() *Gov {
	gov := &Gov{}
	gov.Key = map[string]string{}
	gov.Key["X-API-Key"] = "1vtlJSvzaaB6bTjJKzyakYnjnxrRzM22Ex3j2SDR"
	return gov
}

func (g *Gov) GetGov(url string) (GOV, error) {
	g.Mutex.Lock()
	defer g.Mutex.Unlock()

	key := "X-API-Key"
	value := g.Key[key]

	h := httputils.NewHTTP()
	h.Header(key, value)

	r, err := h.Get(url)

	var gov GOV
	err = json.Unmarshal(r, &gov)
	if err != nil {
		fmt.Println("error:", err)
	}
	return gov, err
}

func (g *Gov) GetSubcommittees(url string) (Subcommittees, error) {

	g.Mutex.Lock()
	defer g.Mutex.Unlock()

	key := "X-API-Key"
	value := g.Key[key]

	h := httputils.NewHTTP()
	h.Header(key, value)

	r, err := h.Get(url)

	var sub Subcommittees
	err = json.Unmarshal(r, &sub)
	if err != nil {
		return sub, err
	}

	return sub, err

}
