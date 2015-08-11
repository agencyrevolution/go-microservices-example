package store

import . "github.com/agencyrevolution/go-microservices-example/models"

type Engine struct {
	State *State
}

type State struct {
	Repos []*Repo
}

func New() *Engine {
	return &Engine{}
}

func (e *Engine) InitializeState() {
	s := &State{
		Repos: []*Repo{

			&Repo{
				Name:            "frontend.react.revolution",
				Description:     "Frontend app for our automation marketing built by React & Flux",
				StargazersCount: 1,
				RepoOwner: &RepoOwner{
					ID:       1,
					Username: "agencyrevolution",
				},
			},

			&Repo{
				Name:            "workflow.microservice.accountlog",
				Description:     "Log account activities/histories to InfluxDB.",
				StargazersCount: 1,
				RepoOwner: &RepoOwner{
					ID:       1,
					Username: "agencyrevolution",
				},
			},

			&Repo{
				Name:            "workflow.microservice.apiserver",
				Description:     "API Server, used as proxy to other microservices",
				StargazersCount: 1,
				RepoOwner: &RepoOwner{
					ID:       1,
					Username: "agencyrevolution",
				},
			},

			&Repo{
				Name:            "workflow.worker.eventdetector",
				Description:     "Event Detector",
				StargazersCount: 1,
				RepoOwner: &RepoOwner{
					ID:       1,
					Username: "agencyrevolution",
				},
			},

			&Repo{
				Name:            "workflow.worker.email",
				Description:     "Sending email using mandrill API.",
				StargazersCount: 1,
				RepoOwner: &RepoOwner{
					ID:       1,
					Username: "agencyrevolution",
				},
			},

			&Repo{
				Name:            "workflow.microservice.url-checker.golang",
				Description:     "Check avaibility of the input URL",
				StargazersCount: 1,
				RepoOwner: &RepoOwner{
					ID:       1,
					Username: "agencyrevolution",
				},
			},

			&Repo{
				Name:            "workflow.microservice.form",
				Description:     "Form API",
				StargazersCount: 1,
				RepoOwner: &RepoOwner{
					ID:       1,
					Username: "agencyrevolution",
				},
			},

			&Repo{
				Name:            "workflow.rule-engine",
				Description:     "Rule engine library",
				StargazersCount: 1,
				RepoOwner: &RepoOwner{
					ID:       1,
					Username: "agencyrevolution",
				},
			},

			&Repo{
				Name:            "frontend.react.growthanalyzer",
				Description:     "An application that graphs customers' potential future revenue when using our product compared to existing growth rates. Once they have gone through the steps they can fill out a form that sends us back information so we can get in touch with them for a consultation.",
				StargazersCount: 1,
				RepoOwner: &RepoOwner{
					ID:       1,
					Username: "agencyrevolution",
				},
			},

			&Repo{
				Name:            "react",
				Description:     "A declarative, efficient, and flexible JavaScript library for building user interfaces.",
				StargazersCount: 25827,
				RepoOwner: &RepoOwner{
					ID:       69631,
					Username: "facebook",
				},
			},

			&Repo{
				Name:            "AsyncDisplayKit",
				Description:     "Smooth asynchronous user interfaces for iOS apps.",
				StargazersCount: 5402,
				RepoOwner: &RepoOwner{
					ID:       69631,
					Username: "facebook",
				},
			},

			&Repo{
				Name:            "atosl",
				Description:     "A partial replacement for Apple's atos tool for converting addresses within a binary file to symbols.",
				StargazersCount: 180,
				RepoOwner: &RepoOwner{
					ID:       69631,
					Username: "facebook",
				},
			},

			&Repo{
				Name:            "augmented-traffic-control",
				Description:     "Augmented Traffic Control: A tool to simulate network conditions",
				StargazersCount: 2459,
				RepoOwner: &RepoOwner{
					ID:       69631,
					Username: "facebook",
				},
			},

			&Repo{
				Name:            "between-meals",
				Description:     "A library to provide calculations between Chef diffs.",
				StargazersCount: 16,
				RepoOwner: &RepoOwner{
					ID:       69631,
					Username: "facebook",
				},
			},

			&Repo{
				Name:            "bistro",
				Description:     "Bistro is a flexible distributed scheduler, a high-performance framework supporting multiple paradigms while retaining ease of configuration, management, and monitoring.",
				StargazersCount: 7,
				RepoOwner: &RepoOwner{
					ID:       69631,
					Username: "facebook",
				},
			},

			&Repo{
				Name:            "buck",
				Description:     "A build system that encourages the creation of small, reusable modules.",
				StargazersCount: 1528,
				RepoOwner: &RepoOwner{
					ID:       69631,
					Username: "facebook",
				},
			},

			&Repo{
				Name:            "C3D",
				Description:     "C3D is a modified version of BVLC caffe to support 3D ConvNets.",
				StargazersCount: 35,
				RepoOwner: &RepoOwner{
					ID:       69631,
					Username: "facebook",
				},
			},

			&Repo{
				Name:            "caf8teen",
				Description:     "",
				StargazersCount: 23,
				RepoOwner: &RepoOwner{
					ID:       69631,
					Username: "facebook",
				},
			},

			&Repo{
				Name:            "chef-utils",
				Description:     "Utilities related to Chef",
				StargazersCount: 165,
				RepoOwner: &RepoOwner{
					ID:       69631,
					Username: "facebook",
				},
			},

			&Repo{
				Name:            "chisel",
				Description:     "Chisel is a collection of LLDB commands to assist debugging iOS apps.",
				StargazersCount: 3595,
				RepoOwner: &RepoOwner{
					ID:       69631,
					Username: "facebook",
				},
			},

			&Repo{
				Name:            "clang-as-ios-dylib",
				Description:     "A workaround to build iOS dynamic libraries from Xcode.",
				StargazersCount: 69,
				RepoOwner: &RepoOwner{
					ID:       69631,
					Username: "facebook",
				},
			},

			&Repo{
				Name:            "codemod",
				Description:     "Codemod is a tool/library to assist you with large-scale codebase refactors that can be partially automated but still require human oversight and occasional intervention. Codemod was developed at Facebook and released as open source.",
				StargazersCount: 583,
				RepoOwner: &RepoOwner{
					ID:       69631,
					Username: "facebook",
				},
			},

			&Repo{
				Name:            "componentkit",
				Description:     "A React-inspired view framework for iOS.",
				StargazersCount: 2253,
				RepoOwner: &RepoOwner{
					ID:       69631,
					Username: "facebook",
				},
			},

			&Repo{
				Name:            "conceal",
				Description:     "Conceal provides easy Android APIs for performing fast encryption and authentication of data.",
				StargazersCount: 1317,
				RepoOwner: &RepoOwner{
					ID:       69631,
					Username: "facebook",
				},
			},

			&Repo{
				Name:            "CParser",
				Description:     "A compact C preprocessor and declaration parser written in pure Lua",
				StargazersCount: 17,
				RepoOwner: &RepoOwner{
					ID:       69631,
					Username: "facebook",
				},
			},

			&Repo{
				Name:            "css-layout",
				Description:     "Reimplementation of CSS layout using pure JavaScript",
				StargazersCount: 2748,
				RepoOwner: &RepoOwner{
					ID:       69631,
					Username: "facebook",
				},
			},

			&Repo{
				Name:            "device-year-class",
				Description:     "A library that analyzes an Android device's specifications and calculates which year the device would be considered \"high end”.",
				StargazersCount: 832,
				RepoOwner: &RepoOwner{
					ID:       69631,
					Username: "facebook",
				},
			},

			&Repo{
				Name:            "dfuse",
				Description:     "dfuse is a binding for the fuse filesystem library written in D.",
				StargazersCount: 29,
				RepoOwner: &RepoOwner{
					ID:       69631,
					Username: "facebook",
				},
			},

			&Repo{
				Name:            "ds2",
				Description:     "Debug server for lldb.",
				StargazersCount: 69,
				RepoOwner: &RepoOwner{
					ID:       69631,
					Username: "facebook",
				},
			},

			&Repo{
				Name:            "easymock",
				Description:     "EasyMock is a Java library that provides an easy way to use Mock Objects in unit testing.",
				StargazersCount: 22,
				RepoOwner: &RepoOwner{
					ID:       69631,
					Username: "facebook",
				},
			},

			&Repo{
				Name:            "emitter",
				Description:     "A JS EventEmitter foundation for evented code",
				StargazersCount: 148,
				RepoOwner: &RepoOwner{
					ID:       69631,
					Username: "facebook",
				},
			},

			&Repo{
				Name:            "esprima",
				Description:     "ECMAScript parsing infrastructure for multipurpose analysis",
				StargazersCount: 93,
				RepoOwner: &RepoOwner{
					ID:       69631,
					Username: "facebook",
				},
			},

			&Repo{
				Name:            "eyescream",
				Description:     "natural image generation using ConvNets",
				StargazersCount: 116,
				RepoOwner: &RepoOwner{
					ID:       69631,
					Username: "facebook",
				},
			},

			&Repo{
				Name:            "facebook-android-sdk",
				Description:     "Used to integrate Android apps with Facebook Platform.",
				StargazersCount: 3391,
				RepoOwner: &RepoOwner{
					ID:       69631,
					Username: "facebook",
				},
			},

			&Repo{
				Name:            "facebook-clang-plugins",
				Description:     "Plugins to clang-analyzer and clang-frontend",
				StargazersCount: 152,
				RepoOwner: &RepoOwner{
					ID:       69631,
					Username: "facebook",
				},
			},

			&Repo{
				Name:            "facebook-ios-sdk",
				Description:     "Used to integrate iOS apps with Facebook Platform.",
				StargazersCount: 4976,
				RepoOwner: &RepoOwner{
					ID:       69631,
					Username: "facebook",
				},
			},

			&Repo{
				Name:            "facebook-oss-pom",
				Description:     "Base POM for facebook open source projects deployed to oss.sonatype.org",
				StargazersCount: 13,
				RepoOwner: &RepoOwner{
					ID:       69631,
					Username: "facebook",
				},
			},

			&Repo{
				Name:            "facebook-php-ads-sdk",
				Description:     "An SDK built to facilitate application development for Facebook Ads API.",
				StargazersCount: 111,
				RepoOwner: &RepoOwner{
					ID:       69631,
					Username: "facebook",
				},
			},

			&Repo{
				Name:            "facebook-php-sdk-v4",
				Description:     "The Facebook SDK for PHP provides a native interface to the Graph API and Facebook Login.  https://developers.facebook.com/docs/php",
				StargazersCount: 813,
				RepoOwner: &RepoOwner{
					ID:       69631,
					Username: "facebook",
				},
			},

			&Repo{
				Name:            "facebook-python-ads-sdk",
				Description:     "An SDK built to facilitate application development for Facebook Ads API.",
				StargazersCount: 109,
				RepoOwner: &RepoOwner{
					ID:       69631,
					Username: "facebook",
				},
			},
		},
	}

	e.State = s
}
