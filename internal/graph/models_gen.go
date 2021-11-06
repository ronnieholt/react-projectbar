// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graph

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/jordanknott/taskcafe/internal/db"
)

type AddTaskLabelInput struct {
	TaskID         uuid.UUID `json:"taskID"`
	ProjectLabelID uuid.UUID `json:"projectLabelID"`
}

type AssignTaskInput struct {
	TaskID uuid.UUID `json:"taskID"`
	UserID uuid.UUID `json:"userID"`
}

type CausedBy struct {
	ID          uuid.UUID    `json:"id"`
	FullName    string       `json:"fullName"`
	ProfileIcon *ProfileIcon `json:"profileIcon"`
}

type ChecklistBadge struct {
	Complete int `json:"complete"`
	Total    int `json:"total"`
}

type CommentsBadge struct {
	Total  int  `json:"total"`
	Unread bool `json:"unread"`
}

type CreateTaskChecklist struct {
	TaskID   uuid.UUID `json:"taskID"`
	Name     string    `json:"name"`
	Position float64   `json:"position"`
}

type CreateTaskChecklistItem struct {
	TaskChecklistID uuid.UUID `json:"taskChecklistID"`
	Name            string    `json:"name"`
	Position        float64   `json:"position"`
}

type CreateTaskComment struct {
	TaskID  uuid.UUID `json:"taskID"`
	Message string    `json:"message"`
}

type CreateTaskCommentPayload struct {
	TaskID  uuid.UUID       `json:"taskID"`
	Comment *db.TaskComment `json:"comment"`
}

type CreateTaskDueDateNotification struct {
	TaskID   uuid.UUID                   `json:"taskID"`
	Period   int                         `json:"period"`
	Duration DueDateNotificationDuration `json:"duration"`
}

type CreateTaskDueDateNotificationsResult struct {
	Notifications []DueDateNotification `json:"notifications"`
}

type CreateTeamMember struct {
	UserID uuid.UUID `json:"userID"`
	TeamID uuid.UUID `json:"teamID"`
}

type CreateTeamMemberPayload struct {
	Team       *db.Team `json:"team"`
	TeamMember *Member  `json:"teamMember"`
}

type CreatedBy struct {
	ID          uuid.UUID    `json:"id"`
	FullName    string       `json:"fullName"`
	ProfileIcon *ProfileIcon `json:"profileIcon"`
}

type DeleteInvitedProjectMember struct {
	ProjectID uuid.UUID `json:"projectID"`
	Email     string    `json:"email"`
}

type DeleteInvitedProjectMemberPayload struct {
	InvitedMember *InvitedMember `json:"invitedMember"`
}

type DeleteInvitedUserAccount struct {
	InvitedUserID uuid.UUID `json:"invitedUserID"`
}

type DeleteInvitedUserAccountPayload struct {
	InvitedUser *InvitedUserAccount `json:"invitedUser"`
}

type DeleteProject struct {
	ProjectID uuid.UUID `json:"projectID"`
}

type DeleteProjectLabel struct {
	ProjectLabelID uuid.UUID `json:"projectLabelID"`
}

type DeleteProjectMember struct {
	ProjectID uuid.UUID `json:"projectID"`
	UserID    uuid.UUID `json:"userID"`
}

type DeleteProjectMemberPayload struct {
	Ok        bool      `json:"ok"`
	Member    *Member   `json:"member"`
	ProjectID uuid.UUID `json:"projectID"`
}

type DeleteProjectPayload struct {
	Ok      bool        `json:"ok"`
	Project *db.Project `json:"project"`
}

type DeleteTaskChecklist struct {
	TaskChecklistID uuid.UUID `json:"taskChecklistID"`
}

type DeleteTaskChecklistItem struct {
	TaskChecklistItemID uuid.UUID `json:"taskChecklistItemID"`
}

type DeleteTaskChecklistItemPayload struct {
	Ok                bool                  `json:"ok"`
	TaskChecklistItem *db.TaskChecklistItem `json:"taskChecklistItem"`
}

type DeleteTaskChecklistPayload struct {
	Ok            bool              `json:"ok"`
	TaskChecklist *db.TaskChecklist `json:"taskChecklist"`
}

type DeleteTaskComment struct {
	CommentID uuid.UUID `json:"commentID"`
}

type DeleteTaskCommentPayload struct {
	TaskID    uuid.UUID `json:"taskID"`
	CommentID uuid.UUID `json:"commentID"`
}

type DeleteTaskDueDateNotification struct {
	ID uuid.UUID `json:"id"`
}

type DeleteTaskDueDateNotificationsResult struct {
	Notifications []uuid.UUID `json:"notifications"`
}

type DeleteTaskGroupInput struct {
	TaskGroupID uuid.UUID `json:"taskGroupID"`
}

type DeleteTaskGroupPayload struct {
	Ok           bool          `json:"ok"`
	AffectedRows int           `json:"affectedRows"`
	TaskGroup    *db.TaskGroup `json:"taskGroup"`
}

type DeleteTaskGroupTasks struct {
	TaskGroupID uuid.UUID `json:"taskGroupID"`
}

type DeleteTaskGroupTasksPayload struct {
	TaskGroupID uuid.UUID   `json:"taskGroupID"`
	Tasks       []uuid.UUID `json:"tasks"`
}

type DeleteTaskInput struct {
	TaskID uuid.UUID `json:"taskID"`
}

type DeleteTaskPayload struct {
	TaskID uuid.UUID `json:"taskID"`
}

type DeleteTeam struct {
	TeamID uuid.UUID `json:"teamID"`
}

type DeleteTeamMember struct {
	TeamID     uuid.UUID  `json:"teamID"`
	UserID     uuid.UUID  `json:"userID"`
	NewOwnerID *uuid.UUID `json:"newOwnerID"`
}

type DeleteTeamMemberPayload struct {
	TeamID           uuid.UUID    `json:"teamID"`
	UserID           uuid.UUID    `json:"userID"`
	AffectedProjects []db.Project `json:"affectedProjects"`
}

type DeleteTeamPayload struct {
	Ok       bool         `json:"ok"`
	Team     *db.Team     `json:"team"`
	Projects []db.Project `json:"projects"`
}

type DeleteUserAccount struct {
	UserID     uuid.UUID  `json:"userID"`
	NewOwnerID *uuid.UUID `json:"newOwnerID"`
}

type DeleteUserAccountPayload struct {
	Ok          bool            `json:"ok"`
	UserAccount *db.UserAccount `json:"userAccount"`
}

type DueDate struct {
	At            *time.Time            `json:"at"`
	Notifications []DueDateNotification `json:"notifications"`
}

type DueDateNotification struct {
	ID       uuid.UUID                   `json:"id"`
	Period   int                         `json:"period"`
	Duration DueDateNotificationDuration `json:"duration"`
}

type DuplicateTaskGroup struct {
	ProjectID   uuid.UUID `json:"projectID"`
	TaskGroupID uuid.UUID `json:"taskGroupID"`
	Name        string    `json:"name"`
	Position    float64   `json:"position"`
}

type DuplicateTaskGroupPayload struct {
	TaskGroup *db.TaskGroup `json:"taskGroup"`
}

type FindProject struct {
	ProjectID      *uuid.UUID `json:"projectID"`
	ProjectShortID *string    `json:"projectShortID"`
}

type FindTask struct {
	TaskID      *uuid.UUID `json:"taskID"`
	TaskShortID *string    `json:"taskShortID"`
}

type FindTeam struct {
	TeamID uuid.UUID `json:"teamID"`
}

type FindUser struct {
	UserID uuid.UUID `json:"userID"`
}

type HasUnreadNotificationsResult struct {
	Unread bool `json:"unread"`
}

type InviteProjectMembers struct {
	ProjectID uuid.UUID      `json:"projectID"`
	Members   []MemberInvite `json:"members"`
}

type InviteProjectMembersPayload struct {
	Ok             bool            `json:"ok"`
	ProjectID      uuid.UUID       `json:"projectID"`
	Members        []Member        `json:"members"`
	InvitedMembers []InvitedMember `json:"invitedMembers"`
}

type InvitedMember struct {
	Email     string    `json:"email"`
	InvitedOn time.Time `json:"invitedOn"`
}

type InvitedUserAccount struct {
	ID        uuid.UUID   `json:"id"`
	Email     string      `json:"email"`
	InvitedOn time.Time   `json:"invitedOn"`
	Member    *MemberList `json:"member"`
}

type LogoutUser struct {
	UserID uuid.UUID `json:"userID"`
}

type MePayload struct {
	User         *db.UserAccount `json:"user"`
	Organization *RoleCode       `json:"organization"`
	TeamRoles    []TeamRole      `json:"teamRoles"`
	ProjectRoles []ProjectRole   `json:"projectRoles"`
}

type Member struct {
	ID          uuid.UUID    `json:"id"`
	Role        *db.Role     `json:"role"`
	FullName    string       `json:"fullName"`
	Username    string       `json:"username"`
	ProfileIcon *ProfileIcon `json:"profileIcon"`
	Owned       *OwnedList   `json:"owned"`
	Member      *MemberList  `json:"member"`
}

type MemberInvite struct {
	UserID *uuid.UUID `json:"userID"`
	Email  *string    `json:"email"`
}

type MemberList struct {
	Teams    []db.Team    `json:"teams"`
	Projects []db.Project `json:"projects"`
}

type MemberSearchFilter struct {
	SearchFilter string     `json:"searchFilter"`
	ProjectID    *uuid.UUID `json:"projectID"`
}

type MemberSearchResult struct {
	Similarity int             `json:"similarity"`
	ID         string          `json:"id"`
	User       *db.UserAccount `json:"user"`
	Status     ShareStatus     `json:"status"`
}

type MyTasks struct {
	Status MyTasksStatus `json:"status"`
	Sort   MyTasksSort   `json:"sort"`
}

type MyTasksPayload struct {
	Tasks    []db.Task            `json:"tasks"`
	Projects []ProjectTaskMapping `json:"projects"`
}

type NewProject struct {
	TeamID *uuid.UUID `json:"teamID"`
	Name   string     `json:"name"`
}

type NewProjectLabel struct {
	ProjectID    uuid.UUID `json:"projectID"`
	LabelColorID uuid.UUID `json:"labelColorID"`
	Name         *string   `json:"name"`
}

type NewTask struct {
	TaskGroupID uuid.UUID   `json:"taskGroupID"`
	Name        string      `json:"name"`
	Position    float64     `json:"position"`
	Assigned    []uuid.UUID `json:"assigned"`
}

type NewTaskGroup struct {
	ProjectID uuid.UUID `json:"projectID"`
	Name      string    `json:"name"`
	Position  float64   `json:"position"`
}

type NewTaskGroupLocation struct {
	TaskGroupID uuid.UUID `json:"taskGroupID"`
	Position    float64   `json:"position"`
}

type NewTaskLocation struct {
	TaskID      uuid.UUID `json:"taskID"`
	TaskGroupID uuid.UUID `json:"taskGroupID"`
	Position    float64   `json:"position"`
}

type NewTeam struct {
	Name           string    `json:"name"`
	OrganizationID uuid.UUID `json:"organizationID"`
}

type NewUserAccount struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	FullName string `json:"fullName"`
	Initials string `json:"initials"`
	Password string `json:"password"`
	RoleCode string `json:"roleCode"`
}

type NotificationCausedBy struct {
	Fullname string    `json:"fullname"`
	Username string    `json:"username"`
	ID       uuid.UUID `json:"id"`
}

type NotificationData struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type NotificationToggleReadInput struct {
	NotifiedID uuid.UUID `json:"notifiedID"`
}

type Notified struct {
	ID           uuid.UUID        `json:"id"`
	Notification *db.Notification `json:"notification"`
	Read         bool             `json:"read"`
	ReadAt       *time.Time       `json:"readAt"`
}

type NotifiedInput struct {
	Limit  int                `json:"limit"`
	Cursor *string            `json:"cursor"`
	Filter NotificationFilter `json:"filter"`
}

type NotifiedResult struct {
	TotalCount int        `json:"totalCount"`
	Notified   []Notified `json:"notified"`
	PageInfo   *PageInfo  `json:"pageInfo"`
}

type OwnedList struct {
	Teams    []db.Team    `json:"teams"`
	Projects []db.Project `json:"projects"`
}

type OwnersList struct {
	Projects []uuid.UUID `json:"projects"`
	Teams    []uuid.UUID `json:"teams"`
}

type PageInfo struct {
	EndCursor   *string `json:"endCursor"`
	HasNextPage bool    `json:"hasNextPage"`
}

type ProfileIcon struct {
	URL      *string `json:"url"`
	Initials *string `json:"initials"`
	BgColor  *string `json:"bgColor"`
}

type ProjectPermission struct {
	Team    RoleCode `json:"team"`
	Project RoleCode `json:"project"`
	Org     RoleCode `json:"org"`
}

type ProjectRole struct {
	ProjectID uuid.UUID `json:"projectID"`
	RoleCode  RoleCode  `json:"roleCode"`
}

type ProjectTaskMapping struct {
	ProjectID uuid.UUID `json:"projectID"`
	TaskID    uuid.UUID `json:"taskID"`
}

type ProjectsFilter struct {
	TeamID *uuid.UUID `json:"teamID"`
}

type RemoveTaskLabelInput struct {
	TaskID      uuid.UUID `json:"taskID"`
	TaskLabelID uuid.UUID `json:"taskLabelID"`
}

type SetTaskChecklistItemComplete struct {
	TaskChecklistItemID uuid.UUID `json:"taskChecklistItemID"`
	Complete            bool      `json:"complete"`
}

type SetTaskComplete struct {
	TaskID   uuid.UUID `json:"taskID"`
	Complete bool      `json:"complete"`
}

type SortTaskGroup struct {
	TaskGroupID uuid.UUID            `json:"taskGroupID"`
	Tasks       []TaskPositionUpdate `json:"tasks"`
}

type SortTaskGroupPayload struct {
	TaskGroupID uuid.UUID `json:"taskGroupID"`
	Tasks       []db.Task `json:"tasks"`
}

type TaskActivityData struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type TaskBadges struct {
	Checklist *ChecklistBadge `json:"checklist"`
	Comments  *CommentsBadge  `json:"comments"`
}

type TaskPositionUpdate struct {
	TaskID   uuid.UUID `json:"taskID"`
	Position float64   `json:"position"`
}

type TeamPermission struct {
	Team RoleCode `json:"team"`
	Org  RoleCode `json:"org"`
}

type TeamRole struct {
	TeamID   uuid.UUID `json:"teamID"`
	RoleCode RoleCode  `json:"roleCode"`
}

type ToggleProjectVisibility struct {
	ProjectID uuid.UUID `json:"projectID"`
	IsPublic  bool      `json:"isPublic"`
}

type ToggleProjectVisibilityPayload struct {
	Project *db.Project `json:"project"`
}

type ToggleTaskLabelInput struct {
	TaskID         uuid.UUID `json:"taskID"`
	ProjectLabelID uuid.UUID `json:"projectLabelID"`
}

type ToggleTaskLabelPayload struct {
	Active bool     `json:"active"`
	Task   *db.Task `json:"task"`
}

type ToggleTaskWatch struct {
	TaskID uuid.UUID `json:"taskID"`
}

type UnassignTaskInput struct {
	TaskID uuid.UUID `json:"taskID"`
	UserID uuid.UUID `json:"userID"`
}

type UpdateProjectLabel struct {
	ProjectLabelID uuid.UUID `json:"projectLabelID"`
	LabelColorID   uuid.UUID `json:"labelColorID"`
	Name           string    `json:"name"`
}

type UpdateProjectLabelColor struct {
	ProjectLabelID uuid.UUID `json:"projectLabelID"`
	LabelColorID   uuid.UUID `json:"labelColorID"`
}

type UpdateProjectLabelName struct {
	ProjectLabelID uuid.UUID `json:"projectLabelID"`
	Name           string    `json:"name"`
}

type UpdateProjectMemberRole struct {
	ProjectID uuid.UUID `json:"projectID"`
	UserID    uuid.UUID `json:"userID"`
	RoleCode  RoleCode  `json:"roleCode"`
}

type UpdateProjectMemberRolePayload struct {
	Ok     bool    `json:"ok"`
	Member *Member `json:"member"`
}

type UpdateProjectName struct {
	ProjectID uuid.UUID `json:"projectID"`
	Name      string    `json:"name"`
}

type UpdateTaskChecklistItemLocation struct {
	TaskChecklistID     uuid.UUID `json:"taskChecklistID"`
	TaskChecklistItemID uuid.UUID `json:"taskChecklistItemID"`
	Position            float64   `json:"position"`
}

type UpdateTaskChecklistItemLocationPayload struct {
	TaskChecklistID uuid.UUID             `json:"taskChecklistID"`
	PrevChecklistID uuid.UUID             `json:"prevChecklistID"`
	ChecklistItem   *db.TaskChecklistItem `json:"checklistItem"`
}

type UpdateTaskChecklistItemName struct {
	TaskChecklistItemID uuid.UUID `json:"taskChecklistItemID"`
	Name                string    `json:"name"`
}

type UpdateTaskChecklistLocation struct {
	TaskChecklistID uuid.UUID `json:"taskChecklistID"`
	Position        float64   `json:"position"`
}

type UpdateTaskChecklistLocationPayload struct {
	Checklist *db.TaskChecklist `json:"checklist"`
}

type UpdateTaskChecklistName struct {
	TaskChecklistID uuid.UUID `json:"taskChecklistID"`
	Name            string    `json:"name"`
}

type UpdateTaskComment struct {
	CommentID uuid.UUID `json:"commentID"`
	Message   string    `json:"message"`
}

type UpdateTaskCommentPayload struct {
	TaskID  uuid.UUID       `json:"taskID"`
	Comment *db.TaskComment `json:"comment"`
}

type UpdateTaskDescriptionInput struct {
	TaskID      uuid.UUID `json:"taskID"`
	Description string    `json:"description"`
}

type UpdateTaskDueDate struct {
	TaskID  uuid.UUID  `json:"taskID"`
	HasTime bool       `json:"hasTime"`
	DueDate *time.Time `json:"dueDate"`
}

type UpdateTaskDueDateNotification struct {
	ID       uuid.UUID                   `json:"id"`
	Period   int                         `json:"period"`
	Duration DueDateNotificationDuration `json:"duration"`
}

type UpdateTaskDueDateNotificationsResult struct {
	Notifications []DueDateNotification `json:"notifications"`
}

type UpdateTaskGroupName struct {
	TaskGroupID uuid.UUID `json:"taskGroupID"`
	Name        string    `json:"name"`
}

type UpdateTaskLocationPayload struct {
	PreviousTaskGroupID uuid.UUID `json:"previousTaskGroupID"`
	Task                *db.Task  `json:"task"`
}

type UpdateTaskName struct {
	TaskID uuid.UUID `json:"taskID"`
	Name   string    `json:"name"`
}

type UpdateTeamMemberRole struct {
	TeamID   uuid.UUID `json:"teamID"`
	UserID   uuid.UUID `json:"userID"`
	RoleCode RoleCode  `json:"roleCode"`
}

type UpdateTeamMemberRolePayload struct {
	Ok     bool      `json:"ok"`
	TeamID uuid.UUID `json:"teamID"`
	Member *Member   `json:"member"`
}

type UpdateUserInfo struct {
	Name     string `json:"name"`
	Initials string `json:"initials"`
	Email    string `json:"email"`
	Bio      string `json:"bio"`
}

type UpdateUserInfoPayload struct {
	User *db.UserAccount `json:"user"`
}

type UpdateUserPassword struct {
	UserID   uuid.UUID `json:"userID"`
	Password string    `json:"password"`
}

type UpdateUserPasswordPayload struct {
	Ok   bool            `json:"ok"`
	User *db.UserAccount `json:"user"`
}

type UpdateUserRole struct {
	UserID   uuid.UUID `json:"userID"`
	RoleCode RoleCode  `json:"roleCode"`
}

type UpdateUserRolePayload struct {
	User *db.UserAccount `json:"user"`
}

type ActionLevel string

const (
	ActionLevelOrg     ActionLevel = "ORG"
	ActionLevelTeam    ActionLevel = "TEAM"
	ActionLevelProject ActionLevel = "PROJECT"
)

var AllActionLevel = []ActionLevel{
	ActionLevelOrg,
	ActionLevelTeam,
	ActionLevelProject,
}

func (e ActionLevel) IsValid() bool {
	switch e {
	case ActionLevelOrg, ActionLevelTeam, ActionLevelProject:
		return true
	}
	return false
}

func (e ActionLevel) String() string {
	return string(e)
}

func (e *ActionLevel) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ActionLevel(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ActionLevel", str)
	}
	return nil
}

func (e ActionLevel) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ActionType string

const (
	ActionTypeTeamAdded              ActionType = "TEAM_ADDED"
	ActionTypeTeamRemoved            ActionType = "TEAM_REMOVED"
	ActionTypeProjectAdded           ActionType = "PROJECT_ADDED"
	ActionTypeProjectRemoved         ActionType = "PROJECT_REMOVED"
	ActionTypeProjectArchived        ActionType = "PROJECT_ARCHIVED"
	ActionTypeDueDateAdded           ActionType = "DUE_DATE_ADDED"
	ActionTypeDueDateRemoved         ActionType = "DUE_DATE_REMOVED"
	ActionTypeDueDateChanged         ActionType = "DUE_DATE_CHANGED"
	ActionTypeTaskAssigned           ActionType = "TASK_ASSIGNED"
	ActionTypeTaskMoved              ActionType = "TASK_MOVED"
	ActionTypeTaskArchived           ActionType = "TASK_ARCHIVED"
	ActionTypeTaskAttachmentUploaded ActionType = "TASK_ATTACHMENT_UPLOADED"
	ActionTypeCommentMentioned       ActionType = "COMMENT_MENTIONED"
	ActionTypeCommentOther           ActionType = "COMMENT_OTHER"
)

var AllActionType = []ActionType{
	ActionTypeTeamAdded,
	ActionTypeTeamRemoved,
	ActionTypeProjectAdded,
	ActionTypeProjectRemoved,
	ActionTypeProjectArchived,
	ActionTypeDueDateAdded,
	ActionTypeDueDateRemoved,
	ActionTypeDueDateChanged,
	ActionTypeTaskAssigned,
	ActionTypeTaskMoved,
	ActionTypeTaskArchived,
	ActionTypeTaskAttachmentUploaded,
	ActionTypeCommentMentioned,
	ActionTypeCommentOther,
}

func (e ActionType) IsValid() bool {
	switch e {
	case ActionTypeTeamAdded, ActionTypeTeamRemoved, ActionTypeProjectAdded, ActionTypeProjectRemoved, ActionTypeProjectArchived, ActionTypeDueDateAdded, ActionTypeDueDateRemoved, ActionTypeDueDateChanged, ActionTypeTaskAssigned, ActionTypeTaskMoved, ActionTypeTaskArchived, ActionTypeTaskAttachmentUploaded, ActionTypeCommentMentioned, ActionTypeCommentOther:
		return true
	}
	return false
}

func (e ActionType) String() string {
	return string(e)
}

func (e *ActionType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ActionType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ActionType", str)
	}
	return nil
}

func (e ActionType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ActivityType string

const (
	ActivityTypeTaskAdded            ActivityType = "TASK_ADDED"
	ActivityTypeTaskMoved            ActivityType = "TASK_MOVED"
	ActivityTypeTaskMarkedComplete   ActivityType = "TASK_MARKED_COMPLETE"
	ActivityTypeTaskMarkedIncomplete ActivityType = "TASK_MARKED_INCOMPLETE"
	ActivityTypeTaskDueDateChanged   ActivityType = "TASK_DUE_DATE_CHANGED"
	ActivityTypeTaskDueDateAdded     ActivityType = "TASK_DUE_DATE_ADDED"
	ActivityTypeTaskDueDateRemoved   ActivityType = "TASK_DUE_DATE_REMOVED"
	ActivityTypeTaskChecklistChanged ActivityType = "TASK_CHECKLIST_CHANGED"
	ActivityTypeTaskChecklistAdded   ActivityType = "TASK_CHECKLIST_ADDED"
	ActivityTypeTaskChecklistRemoved ActivityType = "TASK_CHECKLIST_REMOVED"
)

var AllActivityType = []ActivityType{
	ActivityTypeTaskAdded,
	ActivityTypeTaskMoved,
	ActivityTypeTaskMarkedComplete,
	ActivityTypeTaskMarkedIncomplete,
	ActivityTypeTaskDueDateChanged,
	ActivityTypeTaskDueDateAdded,
	ActivityTypeTaskDueDateRemoved,
	ActivityTypeTaskChecklistChanged,
	ActivityTypeTaskChecklistAdded,
	ActivityTypeTaskChecklistRemoved,
}

func (e ActivityType) IsValid() bool {
	switch e {
	case ActivityTypeTaskAdded, ActivityTypeTaskMoved, ActivityTypeTaskMarkedComplete, ActivityTypeTaskMarkedIncomplete, ActivityTypeTaskDueDateChanged, ActivityTypeTaskDueDateAdded, ActivityTypeTaskDueDateRemoved, ActivityTypeTaskChecklistChanged, ActivityTypeTaskChecklistAdded, ActivityTypeTaskChecklistRemoved:
		return true
	}
	return false
}

func (e ActivityType) String() string {
	return string(e)
}

func (e *ActivityType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ActivityType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ActivityType", str)
	}
	return nil
}

func (e ActivityType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type DueDateNotificationDuration string

const (
	DueDateNotificationDurationMinute DueDateNotificationDuration = "MINUTE"
	DueDateNotificationDurationHour   DueDateNotificationDuration = "HOUR"
	DueDateNotificationDurationDay    DueDateNotificationDuration = "DAY"
	DueDateNotificationDurationWeek   DueDateNotificationDuration = "WEEK"
)

var AllDueDateNotificationDuration = []DueDateNotificationDuration{
	DueDateNotificationDurationMinute,
	DueDateNotificationDurationHour,
	DueDateNotificationDurationDay,
	DueDateNotificationDurationWeek,
}

func (e DueDateNotificationDuration) IsValid() bool {
	switch e {
	case DueDateNotificationDurationMinute, DueDateNotificationDurationHour, DueDateNotificationDurationDay, DueDateNotificationDurationWeek:
		return true
	}
	return false
}

func (e DueDateNotificationDuration) String() string {
	return string(e)
}

func (e *DueDateNotificationDuration) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = DueDateNotificationDuration(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid DueDateNotificationDuration", str)
	}
	return nil
}

func (e DueDateNotificationDuration) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type MyTasksSort string

const (
	MyTasksSortNone    MyTasksSort = "NONE"
	MyTasksSortProject MyTasksSort = "PROJECT"
	MyTasksSortDueDate MyTasksSort = "DUE_DATE"
)

var AllMyTasksSort = []MyTasksSort{
	MyTasksSortNone,
	MyTasksSortProject,
	MyTasksSortDueDate,
}

func (e MyTasksSort) IsValid() bool {
	switch e {
	case MyTasksSortNone, MyTasksSortProject, MyTasksSortDueDate:
		return true
	}
	return false
}

func (e MyTasksSort) String() string {
	return string(e)
}

func (e *MyTasksSort) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = MyTasksSort(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid MyTasksSort", str)
	}
	return nil
}

func (e MyTasksSort) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type MyTasksStatus string

const (
	MyTasksStatusAll               MyTasksStatus = "ALL"
	MyTasksStatusIncomplete        MyTasksStatus = "INCOMPLETE"
	MyTasksStatusCompleteAll       MyTasksStatus = "COMPLETE_ALL"
	MyTasksStatusCompleteToday     MyTasksStatus = "COMPLETE_TODAY"
	MyTasksStatusCompleteYesterday MyTasksStatus = "COMPLETE_YESTERDAY"
	MyTasksStatusCompleteOneWeek   MyTasksStatus = "COMPLETE_ONE_WEEK"
	MyTasksStatusCompleteTwoWeek   MyTasksStatus = "COMPLETE_TWO_WEEK"
	MyTasksStatusCompleteThreeWeek MyTasksStatus = "COMPLETE_THREE_WEEK"
)

var AllMyTasksStatus = []MyTasksStatus{
	MyTasksStatusAll,
	MyTasksStatusIncomplete,
	MyTasksStatusCompleteAll,
	MyTasksStatusCompleteToday,
	MyTasksStatusCompleteYesterday,
	MyTasksStatusCompleteOneWeek,
	MyTasksStatusCompleteTwoWeek,
	MyTasksStatusCompleteThreeWeek,
}

func (e MyTasksStatus) IsValid() bool {
	switch e {
	case MyTasksStatusAll, MyTasksStatusIncomplete, MyTasksStatusCompleteAll, MyTasksStatusCompleteToday, MyTasksStatusCompleteYesterday, MyTasksStatusCompleteOneWeek, MyTasksStatusCompleteTwoWeek, MyTasksStatusCompleteThreeWeek:
		return true
	}
	return false
}

func (e MyTasksStatus) String() string {
	return string(e)
}

func (e *MyTasksStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = MyTasksStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid MyTasksStatus", str)
	}
	return nil
}

func (e MyTasksStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type NotificationFilter string

const (
	NotificationFilterAll       NotificationFilter = "ALL"
	NotificationFilterUnread    NotificationFilter = "UNREAD"
	NotificationFilterAssigned  NotificationFilter = "ASSIGNED"
	NotificationFilterMentioned NotificationFilter = "MENTIONED"
)

var AllNotificationFilter = []NotificationFilter{
	NotificationFilterAll,
	NotificationFilterUnread,
	NotificationFilterAssigned,
	NotificationFilterMentioned,
}

func (e NotificationFilter) IsValid() bool {
	switch e {
	case NotificationFilterAll, NotificationFilterUnread, NotificationFilterAssigned, NotificationFilterMentioned:
		return true
	}
	return false
}

func (e NotificationFilter) String() string {
	return string(e)
}

func (e *NotificationFilter) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = NotificationFilter(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid NotificationFilter", str)
	}
	return nil
}

func (e NotificationFilter) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ObjectType string

const (
	ObjectTypeOrg               ObjectType = "ORG"
	ObjectTypeTeam              ObjectType = "TEAM"
	ObjectTypeProject           ObjectType = "PROJECT"
	ObjectTypeTask              ObjectType = "TASK"
	ObjectTypeTaskGroup         ObjectType = "TASK_GROUP"
	ObjectTypeTaskChecklist     ObjectType = "TASK_CHECKLIST"
	ObjectTypeTaskChecklistItem ObjectType = "TASK_CHECKLIST_ITEM"
)

var AllObjectType = []ObjectType{
	ObjectTypeOrg,
	ObjectTypeTeam,
	ObjectTypeProject,
	ObjectTypeTask,
	ObjectTypeTaskGroup,
	ObjectTypeTaskChecklist,
	ObjectTypeTaskChecklistItem,
}

func (e ObjectType) IsValid() bool {
	switch e {
	case ObjectTypeOrg, ObjectTypeTeam, ObjectTypeProject, ObjectTypeTask, ObjectTypeTaskGroup, ObjectTypeTaskChecklist, ObjectTypeTaskChecklistItem:
		return true
	}
	return false
}

func (e ObjectType) String() string {
	return string(e)
}

func (e *ObjectType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ObjectType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ObjectType", str)
	}
	return nil
}

func (e ObjectType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type RoleCode string

const (
	RoleCodeOwner    RoleCode = "owner"
	RoleCodeAdmin    RoleCode = "admin"
	RoleCodeMember   RoleCode = "member"
	RoleCodeObserver RoleCode = "observer"
)

var AllRoleCode = []RoleCode{
	RoleCodeOwner,
	RoleCodeAdmin,
	RoleCodeMember,
	RoleCodeObserver,
}

func (e RoleCode) IsValid() bool {
	switch e {
	case RoleCodeOwner, RoleCodeAdmin, RoleCodeMember, RoleCodeObserver:
		return true
	}
	return false
}

func (e RoleCode) String() string {
	return string(e)
}

func (e *RoleCode) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = RoleCode(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid RoleCode", str)
	}
	return nil
}

func (e RoleCode) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type RoleLevel string

const (
	RoleLevelAdmin  RoleLevel = "ADMIN"
	RoleLevelMember RoleLevel = "MEMBER"
)

var AllRoleLevel = []RoleLevel{
	RoleLevelAdmin,
	RoleLevelMember,
}

func (e RoleLevel) IsValid() bool {
	switch e {
	case RoleLevelAdmin, RoleLevelMember:
		return true
	}
	return false
}

func (e RoleLevel) String() string {
	return string(e)
}

func (e *RoleLevel) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = RoleLevel(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid RoleLevel", str)
	}
	return nil
}

func (e RoleLevel) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ShareStatus string

const (
	ShareStatusInvited ShareStatus = "INVITED"
	ShareStatusJoined  ShareStatus = "JOINED"
)

var AllShareStatus = []ShareStatus{
	ShareStatusInvited,
	ShareStatusJoined,
}

func (e ShareStatus) IsValid() bool {
	switch e {
	case ShareStatusInvited, ShareStatusJoined:
		return true
	}
	return false
}

func (e ShareStatus) String() string {
	return string(e)
}

func (e *ShareStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ShareStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ShareStatus", str)
	}
	return nil
}

func (e ShareStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
