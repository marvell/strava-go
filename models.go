package strava

import "time"

// Fault represents a Strava API error response
type Fault struct {
	Errors   []Error `json:"errors"`
	Message  string  `json:"message"`
	Resource string  `json:"resource"`
}

// Error represents a specific error in a Fault response
type Error struct {
	Code     string `json:"code"`
	Field    string `json:"field"`
	Resource string `json:"resource"`
}

// Athlete represents a Strava athlete
type Athlete struct {
	ID                    uint    `json:"id"`
	FirstName             string  `json:"firstname"`
	LastName              string  `json:"lastname"`
	ProfileMedium         string  `json:"profile_medium"`
	Profile               string  `json:"profile"`
	City                  string  `json:"city"`
	State                 string  `json:"state"`
	Country               string  `json:"country"`
	Sex                   string  `json:"sex"`
	Premium               bool    `json:"premium"`
	Summit                bool    `json:"summit"`
	CreatedAt             string  `json:"created_at"`
	UpdatedAt             string  `json:"updated_at"`
	FollowerCount         int     `json:"follower_count"`
	FriendCount           int     `json:"friend_count"`
	MeasurementPreference string  `json:"measurement_preference"`
	FTP                   int     `json:"ftp"`
	Weight                float64 `json:"weight"`
	Clubs                 []Club  `json:"clubs"`
	Bikes                 []Gear  `json:"bikes"`
	Shoes                 []Gear  `json:"shoes"`
}

// DetailedAthlete extends Athlete with additional fields
type DetailedAthlete struct {
	Athlete
	ResourceState     int    `json:"resource_state"`
	MutualFriendCount int    `json:"mutual_friend_count"`
	AthleteType       int    `json:"athlete_type"`
	DatePreference    string `json:"date_preference"`
}

// Club represents a Strava club
type Club struct {
	ID              uint      `json:"id"`
	Name            string    `json:"name"`
	ProfileMedium   string    `json:"profile_medium"`
	Profile         string    `json:"profile"`
	CoverPhoto      string    `json:"cover_photo"`
	CoverPhotoSmall string    `json:"cover_photo_small"`
	SportType       SportType `json:"sport_type"`
	City            string    `json:"city"`
	State           string    `json:"state"`
	Country         string    `json:"country"`
	Private         bool      `json:"private"`
	MemberCount     int       `json:"member_count"`
	Featured        bool      `json:"featured"`
	Verified        bool      `json:"verified"`
	Url             string    `json:"url"`
}

// Gear represents equipment like bikes or shoes
type Gear struct {
	ID            string  `json:"id"`
	Primary       bool    `json:"primary"`
	Name          string  `json:"name"`
	ResourceState int     `json:"resource_state"`
	Distance      float64 `json:"distance"`
	BrandName     string  `json:"brand_name"`
	ModelName     string  `json:"model_name"`
	FrameType     int     `json:"frame_type,omitempty"`
	Description   string  `json:"description"`
}

// Map represents a route map
type Map struct {
	ID              string `json:"id"`
	Polyline        string `json:"polyline"`
	SummaryPolyline string `json:"summary_polyline"`
}

// SummaryActivity represents a summary of an activity on Strava
type SummaryActivity struct {
	ID                   uint         `json:"id"`
	ExternalID           string       `json:"external_id"`
	UploadID             uint         `json:"upload_id"`
	Athlete              *Athlete     `json:"athlete"`
	Name                 string       `json:"name"`
	Distance             float64      `json:"distance"`
	MovingTime           int          `json:"moving_time"`
	ElapsedTime          int          `json:"elapsed_time"`
	TotalElevationGain   float64      `json:"total_elevation_gain"`
	ElevHigh             float64      `json:"elev_high"`
	ElevLow              float64      `json:"elev_low"`
	Type                 ActivityType `json:"type"`
	SportType            SportType    `json:"sport_type"`
	StartDate            time.Time    `json:"start_date"`
	StartDateLocal       time.Time    `json:"start_date_local"`
	Timezone             string       `json:"timezone"`
	StartLatLng          []float64    `json:"start_latlng"`
	EndLatLng            []float64    `json:"end_latlng"`
	AchievementCount     int          `json:"achievement_count"`
	KudosCount           int          `json:"kudos_count"`
	CommentCount         int          `json:"comment_count"`
	AthleteCount         int          `json:"athlete_count"`
	PhotoCount           int          `json:"photo_count"`
	TotalPhotoCount      int          `json:"total_photo_count"`
	Map                  *Map         `json:"map"`
	Trainer              bool         `json:"trainer"`
	Commute              bool         `json:"commute"`
	Manual               bool         `json:"manual"`
	Private              bool         `json:"private"`
	Flagged              bool         `json:"flagged"`
	WorkoutType          int          `json:"workout_type"`
	GearID               string       `json:"gear_id"`
	AverageSpeed         float64      `json:"average_speed"`
	MaxSpeed             float64      `json:"max_speed"`
	AverageCadence       float64      `json:"average_cadence"`
	AverageTemp          float64      `json:"average_temp"`
	AverageWatts         float64      `json:"average_watts"`
	WeightedAverageWatts int          `json:"weighted_average_watts"`
	Kilojoules           float64      `json:"kilojoules"`
	DeviceWatts          bool         `json:"device_watts"`
	HasHeartrate         bool         `json:"has_heartrate"`
	AverageHeartrate     float64      `json:"average_heartrate"`
	MaxHeartrate         float64      `json:"max_heartrate"`
	MaxWatts             int          `json:"max_watts"`
	SufferScore          int          `json:"suffer_score"`
}

// DetailedActivity represents a detailed activity on Strava
type DetailedActivity struct {
	SummaryActivity
	Description    string                   `json:"description"`
	Photos         *PhotosSummary           `json:"photos"`
	Gear           *SummaryGear             `json:"gear"`
	Calories       float64                  `json:"calories"`
	SegmentEfforts []*DetailedSegmentEffort `json:"segment_efforts"`
	DeviceName     string                   `json:"device_name"`
	EmbedToken     string                   `json:"embed_token"`
	SplitsMetric   []*Split                 `json:"splits_metric"`
	SplitsStandard []*Split                 `json:"splits_standard"`
	Laps           []*Lap                   `json:"laps"`
	BestEfforts    []*DetailedSegmentEffort `json:"best_efforts"`
}

// PhotosSummary represents a summary of photos for an activity
type PhotosSummary struct {
	Count   int           `json:"count"`
	Primary *PhotoSummary `json:"primary"`
}

// PhotoSummary represents a summary of a photo
type PhotoSummary struct {
	ID       uint              `json:"id"`
	Source   int               `json:"source"`
	UniqueID string            `json:"unique_id"`
	URLs     map[string]string `json:"urls"`
}

// SummaryGear represents a summary of gear used during an activity
type SummaryGear struct {
	ID            string  `json:"id"`
	Primary       bool    `json:"primary"`
	Name          string  `json:"name"`
	ResourceState int     `json:"resource_state"`
	Distance      float64 `json:"distance"`
}

// DetailedSegmentEffort represents a detailed segment effort
type DetailedSegmentEffort struct {
	ID               uint            `json:"id"`
	ResourceState    int             `json:"resource_state"`
	Name             string          `json:"name"`
	Activity         *MetaActivity   `json:"activity"`
	Athlete          *MetaAthlete    `json:"athlete"`
	ElapsedTime      int             `json:"elapsed_time"`
	MovingTime       int             `json:"moving_time"`
	StartDate        time.Time       `json:"start_date"`
	StartDateLocal   time.Time       `json:"start_date_local"`
	Distance         float64         `json:"distance"`
	StartIndex       int             `json:"start_index"`
	EndIndex         int             `json:"end_index"`
	AverageCadence   float64         `json:"average_cadence"`
	AverageWatts     float64         `json:"average_watts"`
	DeviceWatts      bool            `json:"device_watts"`
	AverageHeartrate float64         `json:"average_heartrate"`
	MaxHeartrate     float64         `json:"max_heartrate"`
	Segment          *SummarySegment `json:"segment"`
	KOMRank          int             `json:"kom_rank"`
	PRRank           int             `json:"pr_rank"`
	Hidden           bool            `json:"hidden"`
}

// Split represents a split in an activity
type Split struct {
	Distance                  float64 `json:"distance"`
	ElapsedTime               int     `json:"elapsed_time"`
	ElevationDifference       float64 `json:"elevation_difference"`
	MovingTime                int     `json:"moving_time"`
	Split                     int     `json:"split"`
	AverageSpeed              float64 `json:"average_speed"`
	AverageGradeAdjustedSpeed float64 `json:"average_grade_adjusted_speed"`
	AverageHeartrate          float64 `json:"average_heartrate"`
	PaceZone                  int     `json:"pace_zone"`
}

// Lap represents a lap in an activity
type Lap struct {
	ID                 uint          `json:"id"`
	ResourceState      int           `json:"resource_state"`
	Name               string        `json:"name"`
	Activity           *MetaActivity `json:"activity"`
	Athlete            *MetaAthlete  `json:"athlete"`
	ElapsedTime        int           `json:"elapsed_time"`
	MovingTime         int           `json:"moving_time"`
	StartDate          time.Time     `json:"start_date"`
	StartDateLocal     time.Time     `json:"start_date_local"`
	Distance           float64       `json:"distance"`
	StartIndex         int           `json:"start_index"`
	EndIndex           int           `json:"end_index"`
	TotalElevationGain float64       `json:"total_elevation_gain"`
	AverageSpeed       float64       `json:"average_speed"`
	MaxSpeed           float64       `json:"max_speed"`
	AverageCadence     float64       `json:"average_cadence"`
	DeviceWatts        bool          `json:"device_watts"`
	AverageWatts       float64       `json:"average_watts"`
	LapIndex           int           `json:"lap_index"`
	Split              int           `json:"split"`
	AverageHeartrate   float64       `json:"average_heartrate"`
}

// MetaActivity represents minimal activity data
type MetaActivity struct {
	ID            uint `json:"id"`
	ResourceState int  `json:"resource_state"`
}

// MetaAthlete represents minimal athlete data
type MetaAthlete struct {
	ID            uint `json:"id"`
	ResourceState int  `json:"resource_state"`
}

// SummarySegment represents a summary of a segment
type SummarySegment struct {
	ID            uint      `json:"id"`
	ResourceState int       `json:"resource_state"`
	Name          string    `json:"name"`
	ActivityType  string    `json:"activity_type"`
	Distance      float64   `json:"distance"`
	AverageGrade  float64   `json:"average_grade"`
	MaximumGrade  float64   `json:"maximum_grade"`
	ElevationHigh float64   `json:"elevation_high"`
	ElevationLow  float64   `json:"elevation_low"`
	StartLatLng   []float64 `json:"start_latlng"`
	EndLatLng     []float64 `json:"end_latlng"`
	ClimbCategory int       `json:"climb_category"`
	City          string    `json:"city"`
	State         string    `json:"state"`
	Country       string    `json:"country"`
	Private       bool      `json:"private"`
	Hazardous     bool      `json:"hazardous"`
	Starred       bool      `json:"starred"`
}

// SportType represents the type of sport/activity
type SportType string

const (
	SportTypeAlpineSki                     SportType = "AlpineSki"
	SportTypeBackcountrySki                SportType = "BackcountrySki"
	SportTypeBadminton                     SportType = "Badminton"
	SportTypeCanoeing                      SportType = "Canoeing"
	SportTypeCrossfit                      SportType = "Crossfit"
	SportTypeEBikeRide                     SportType = "EBikeRide"
	SportTypeElliptical                    SportType = "Elliptical"
	SportTypeEMountainBikeRide             SportType = "EMountainBikeRide"
	SportTypeGolf                          SportType = "Golf"
	SportTypeGravelRide                    SportType = "GravelRide"
	SportTypeHandcycle                     SportType = "Handcycle"
	SportTypeHighIntensityIntervalTraining SportType = "HighIntensityIntervalTraining"
	SportTypeHike                          SportType = "Hike"
	SportTypeIceSkate                      SportType = "IceSkate"
	SportTypeInlineSkate                   SportType = "InlineSkate"
	SportTypeKayaking                      SportType = "Kayaking"
	SportTypeKitesurf                      SportType = "Kitesurf"
	SportTypeMountainBikeRide              SportType = "MountainBikeRide"
	SportTypeNordicSki                     SportType = "NordicSki"
	SportTypePickleball                    SportType = "Pickleball"
	SportTypePilates                       SportType = "Pilates"
	SportTypeRacquetball                   SportType = "Racquetball"
	SportTypeRide                          SportType = "Ride"
	SportTypeRockClimbing                  SportType = "RockClimbing"
	SportTypeRollerSki                     SportType = "RollerSki"
	SportTypeRowing                        SportType = "Rowing"
	SportTypeRun                           SportType = "Run"
	SportTypeSail                          SportType = "Sail"
	SportTypeSkateboard                    SportType = "Skateboard"
	SportTypeSnowboard                     SportType = "Snowboard"
	SportTypeSnowshoe                      SportType = "Snowshoe"
	SportTypeSoccer                        SportType = "Soccer"
	SportTypeSquash                        SportType = "Squash"
	SportTypeStairStepper                  SportType = "StairStepper"
	SportTypeStandUpPaddling               SportType = "StandUpPaddling"
	SportTypeSurfing                       SportType = "Surfing"
	SportTypeSwim                          SportType = "Swim"
	SportTypeTableTennis                   SportType = "TableTennis"
	SportTypeTennis                        SportType = "Tennis"
	SportTypeTrailRun                      SportType = "TrailRun"
	SportTypeVelomobile                    SportType = "Velomobile"
	SportTypeVirtualRide                   SportType = "VirtualRide"
	SportTypeVirtualRow                    SportType = "VirtualRow"
	SportTypeVirtualRun                    SportType = "VirtualRun"
	SportTypeWalk                          SportType = "Walk"
	SportTypeWeightTraining                SportType = "WeightTraining"
	SportTypeWheelchair                    SportType = "Wheelchair"
	SportTypeWindsurf                      SportType = "Windsurf"
	SportTypeWorkout                       SportType = "Workout"
	SportTypeYoga                          SportType = "Yoga"
)

// ActivityType represents the type of activity
type ActivityType string

const (
	ActivityTypeAlpineSki       ActivityType = "AlpineSki"
	ActivityTypeBackcountrySki  ActivityType = "BackcountrySki"
	ActivityTypeCanoeing        ActivityType = "Canoeing"
	ActivityTypeCrossfit        ActivityType = "Crossfit"
	ActivityTypeEBikeRide       ActivityType = "EBikeRide"
	ActivityTypeElliptical      ActivityType = "Elliptical"
	ActivityTypeGolf            ActivityType = "Golf"
	ActivityTypeHandcycle       ActivityType = "Handcycle"
	ActivityTypeHike            ActivityType = "Hike"
	ActivityTypeIceSkate        ActivityType = "IceSkate"
	ActivityTypeInlineSkate     ActivityType = "InlineSkate"
	ActivityTypeKayaking        ActivityType = "Kayaking"
	ActivityTypeKitesurf        ActivityType = "Kitesurf"
	ActivityTypeNordicSki       ActivityType = "NordicSki"
	ActivityTypeRide            ActivityType = "Ride"
	ActivityTypeRockClimbing    ActivityType = "RockClimbing"
	ActivityTypeRollerSki       ActivityType = "RollerSki"
	ActivityTypeRowing          ActivityType = "Rowing"
	ActivityTypeRun             ActivityType = "Run"
	ActivityTypeSail            ActivityType = "Sail"
	ActivityTypeSkateboard      ActivityType = "Skateboard"
	ActivityTypeSnowboard       ActivityType = "Snowboard"
	ActivityTypeSnowshoe        ActivityType = "Snowshoe"
	ActivityTypeSoccer          ActivityType = "Soccer"
	ActivityTypeStairStepper    ActivityType = "StairStepper"
	ActivityTypeStandUpPaddling ActivityType = "StandUpPaddling"
	ActivityTypeSurfing         ActivityType = "Surfing"
	ActivityTypeSwim            ActivityType = "Swim"
	ActivityTypeVelomobile      ActivityType = "Velomobile"
	ActivityTypeVirtualRide     ActivityType = "VirtualRide"
	ActivityTypeVirtualRun      ActivityType = "VirtualRun"
	ActivityTypeWalk            ActivityType = "Walk"
	ActivityTypeWeightTraining  ActivityType = "WeightTraining"
	ActivityTypeWheelchair      ActivityType = "Wheelchair"
	ActivityTypeWindsurf        ActivityType = "Windsurf"
	ActivityTypeWorkout         ActivityType = "Workout"
	ActivityTypeYoga            ActivityType = "Yoga"
)
