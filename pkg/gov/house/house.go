package house

type House struct {
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
			LeadershipRole       string      `json:"leadership_role"`
			TwitterAccount       string      `json:"twitter_account"`
			FacebookAccount      string      `json:"facebook_account"`
			YoutubeAccount       interface{} `json:"youtube_account"`
			GovtrackID           string      `json:"govtrack_id"`
			CspanID              string      `json:"cspan_id"`
			VotesmartID          string      `json:"votesmart_id"`
			IcpsrID              string      `json:"icpsr_id"`
			CrpID                string      `json:"crp_id"`
			GoogleEntityID       string      `json:"google_entity_id"`
			FecCandidateID       string      `json:"fec_candidate_id"`
			URL                  string      `json:"url"`
			RssURL               string      `json:"rss_url"`
			ContactForm          interface{} `json:"contact_form"`
			InOffice             bool        `json:"in_office"`
			CookPvi              string      `json:"cook_pvi"`
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
			Fax                  interface{} `json:"fax"`
			State                string      `json:"state"`
			District             string      `json:"district"`
			AtLarge              bool        `json:"at_large"`
			Geoid                string      `json:"geoid"`
			MissedVotesPct       float64     `json:"missed_votes_pct,omitempty"`
			VotesWithPartyPct    float64     `json:"votes_with_party_pct,omitempty"`
			VotesAgainstPartyPct float64     `json:"votes_against_party_pct,omitempty"`
		} `json:"members"`
	} `json:"results"`
}
