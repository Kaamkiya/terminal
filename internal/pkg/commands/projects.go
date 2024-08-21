package commands

import (
	"fmt"
	"encoding/json"
	"io"
	"time"
	"net/http"

	"codeberg.org/Kaamkiya/terminal/internal/pkg/style"

	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/lipgloss/table"
)

type repo struct {
	ID    int `json:"id"`
	Owner interface{} `json:"owner"`
	Name            string      `json:"name"`
	FullName        string      `json:"full_name"`
	Description     string      `json:"description"`
	Empty           bool        `json:"empty"`
	Private         bool        `json:"private"`
	Fork            bool        `json:"fork"`
	Template        bool        `json:"template"`
	Parent          interface{} `json:"parent"`
	Mirror          bool        `json:"mirror"`
	Size            int         `json:"size"`
	Language        string      `json:"language"`
	LanguagesURL    string      `json:"languages_url"`
	HTMLURL         string      `json:"html_url"`
	URL             string      `json:"url"`
	Link            string      `json:"link"`
	SSHURL          string      `json:"ssh_url"`
	CloneURL        string      `json:"clone_url"`
	OriginalURL     string      `json:"original_url"`
	Website         string      `json:"website"`
	StarsCount      int         `json:"stars_count"`
	ForksCount      int         `json:"forks_count"`
	WatchersCount   int         `json:"watchers_count"`
	OpenIssuesCount int         `json:"open_issues_count"`
	OpenPrCounter   int         `json:"open_pr_counter"`
	ReleaseCounter  int         `json:"release_counter"`
	DefaultBranch   string      `json:"default_branch"`
	Archived        bool        `json:"archived"`
	CreatedAt       time.Time   `json:"created_at"`
	UpdatedAt       time.Time   `json:"updated_at"`
	ArchivedAt      time.Time   `json:"archived_at"`
	Permissions     struct {
		Admin bool `json:"admin"`
		Push  bool `json:"push"`
		Pull  bool `json:"pull"`
	} `json:"permissions"`
	HasIssues       bool `json:"has_issues"`
	InternalTracker struct {
		EnableTimeTracker                bool `json:"enable_time_tracker"`
		AllowOnlyContributorsToTrackTime bool `json:"allow_only_contributors_to_track_time"`
		EnableIssueDependencies          bool `json:"enable_issue_dependencies"`
	} `json:"internal_tracker"`
	HasWiki                       bool        `json:"has_wiki"`
	WikiBranch                    string      `json:"wiki_branch"`
	GloballyEditableWiki          bool        `json:"globally_editable_wiki"`
	HasPullRequests               bool        `json:"has_pull_requests"`
	HasProjects                   bool        `json:"has_projects"`
	HasReleases                   bool        `json:"has_releases"`
	HasPackages                   bool        `json:"has_packages"`
	HasActions                    bool        `json:"has_actions"`
	IgnoreWhitespaceConflicts     bool        `json:"ignore_whitespace_conflicts"`
	AllowMergeCommits             bool        `json:"allow_merge_commits"`
	AllowRebase                   bool        `json:"allow_rebase"`
	AllowRebaseExplicit           bool        `json:"allow_rebase_explicit"`
	AllowSquashMerge              bool        `json:"allow_squash_merge"`
	AllowFastForwardOnlyMerge     bool        `json:"allow_fast_forward_only_merge"`
	AllowRebaseUpdate             bool        `json:"allow_rebase_update"`
	DefaultDeleteBranchAfterMerge bool        `json:"default_delete_branch_after_merge"`
	DefaultMergeStyle             string      `json:"default_merge_style"`
	DefaultAllowMaintainerEdit    bool        `json:"default_allow_maintainer_edit"`
	AvatarURL                     string      `json:"avatar_url"`
	Internal                      bool        `json:"internal"`
	MirrorInterval                string      `json:"mirror_interval"`
	ObjectFormatName              string      `json:"object_format_name"`
	MirrorUpdated                 time.Time   `json:"mirror_updated"`
	RepoTransfer                  interface{} `json:"repo_transfer"`
	Topics                        []string    `json:"topics"`
}

func projectsCmd(session ssh.Session, styles style.Style) {
	apiURL := "https://codeberg.org/api/v1/users/Kaamkiya/repos"

	repos := []repo{}

	res, err := http.Get(apiURL)
	if err != nil {
		fmt.Fprintln(session, "Sorry, there was an error getting the repositories.")
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Fprintln(session, "Sorry, there was an error getting the repositories.")
		return
	}

	err = json.Unmarshal(body, &repos)
	if err != nil {
		fmt.Fprintln(session, "Sorry, there was an error getting the repositories.")
		return
	}

	headers := []string{"Name", "Last update", "Description", "Language", "URL"}
	data := [][]string{}

	for _, repo := range repos {
		data = append(data, []string{
			styles.Green.Render(repo.FullName),
			repo.UpdatedAt.Format("2006/02/01"),
			repo.Description,
			repo.Language,
			repo.HTMLURL,
		})
	}

	t := table.New().
		Headers(headers...).
		Rows(data...).
		Render()

	fmt.Fprintln(session, t)
}
