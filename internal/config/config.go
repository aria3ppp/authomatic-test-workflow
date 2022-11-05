package config

import (
	"errors"
	"log"
	"os"
	"path/filepath"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

var Config config

func Load(configFilename string) error {
	err := godotenv.Load(filepath.Join(filepath.Dir(configFilename), ".env"))
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return err
		}
		log.Println(err.Error())
	}
	return cleanenv.ReadConfig(configFilename, &Config)
}

type config struct {
	Servic struct {
		Database struct {
			DSN string `yaml:"dsn" env:"DSN" env-required:"true"`
		} `yaml:"database" env-required:"true"`

		Server struct {
			Production               bool   `yaml:"production" env:"SERVER_PRODUCTION" env-default:"false"`
			Logfile                  string `yaml:"logfile" env:"SERVER_LOGFILE" env-required:"true"`
			Port                     uint16 `yaml:"port" env:"SERVER_PORT" env-required:"true"`
			HandlerTimeoutInSeconds  int    `yaml:"handler_timeout_in_seconds" env-required:"true"`
			ShutdownTimeoutInSeconds int    `yaml:"shutdown_timeout_in_seconds" env-required:"true"`
		} `yaml:"server" env-required:"true"`

		Token struct {
			SecretKey string `yaml:"secret_key" env:"SERVER_SECRET_KEY" env-required:"true"`
			Access    struct {
				Duration struct {
					InMinutes int `yaml:"in_minutes" env-required:"true"`
				} `yaml:"duration" env-required:"true"`
			} `yaml:"access" env-required:"true"`
			Refresh struct {
				Duration struct {
					InMinutes int `yaml:"in_minutes" env-required:"true"`
				} `yaml:"duration" env-required:"true"`
			} `yaml:"refresh" env-required:"true"`
		} `yaml:"token" env-required:"true"`

		Elasticsearch struct {
			Url string `yaml:"url" env:"ELASTICSEARCH_URL" env-required:"true"`
		} `yaml:"elasticsearch" env-required:"true"`
	} `yaml:"service" env-required:"true"`

	Pagination struct {
		Page struct {
			VarName  string `yaml:"var_name" env-required:"true"`
			MinValue int    `yaml:"min_value" env-required:"true"`
		} `yaml:"page" env-required:"true"`
		PageSize struct {
			VarName      string `yaml:"var_name" env-required:"true"`
			DefaultValue int    `yaml:"default_value" env-required:"true"`
			MinValue     int    `yaml:"min_value" env-required:"true"`
			MaxValue     int    `yaml:"max_value" env-required:"true"`
		} `yaml:"page_size" env-required:"true"`
	} `yaml:"pagination" env-required:"true"`

	Validation struct {
		Request struct {
			Search struct {
				Query struct {
					MinLength int `yaml:"min_length" env-required:"true"`
					MaxLength int `yaml:"max_length" env-required:"true"`
				} `yaml:"query" env-required:"true"`
			} `yaml:"search" env-required:"true"`
			Invalidation struct {
				MinLength int `yaml:"min_length" env-required:"true"`
				MaxLength int `yaml:"max_length" env-required:"true"`
			} `yaml:"invalidation" env-required:"true"`
			Array struct {
				MaxLength int `yaml:"max_length" env-required:"true"`
			} `yaml:"array" env-required:"true"`
		} `yaml:"request" env-required:"true"`

		User struct {
			Email struct {
				MinLength int `yaml:"min_length" env-required:"true"`
				MaxLength int `yaml:"max_length" env-required:"true"`
			} `yaml:"email" env-required:"true"`
			Password struct {
				MinLength            int `yaml:"min_length" env-required:"true"`
				MaxLength            int `yaml:"max_length" env-required:"true"`
				RequiredNumbers      int `yaml:"required_numbers" env-required:"true"`
				RequiredLowerLetters int `yaml:"required_lower_letters" env-required:"true"`
				RequiredUpperLetters int `yaml:"required_upper_letters" env-required:"true"`
				RequiredSpecialChars int `yaml:"required_special_chars" env-required:"true"`
			} `yaml:"password" env-required:"true"`
			FirstName struct {
				MinLength int `yaml:"min_length" env-required:"true"`
				MaxLength int `yaml:"max_length" env-required:"true"`
			} `yaml:"first_name" env-required:"true"`
			LastName struct {
				MinLength int `yaml:"min_length" env-required:"true"`
				MaxLength int `yaml:"max_length" env-required:"true"`
			} `yaml:"last_name" env-required:"true"`
			Bio struct {
				MinLength int `yaml:"min_length" env-required:"true"`
				MaxLength int `yaml:"max_length" env-required:"true"`
			} `yaml:"bio" env-required:"true"`
			Birthdate struct {
				MinValue struct {
					Year  int `yaml:"year"  env-required:"true"`
					Month int `yaml:"month"  env-required:"true"`
					Day   int `yaml:"day"  env-required:"true"`
				} `yaml:"min_value" env-required:"true"`
			} `yaml:"birthdate" env-required:"true"`
		} `yaml:"user" env-required:"true"`

		Film struct {
			Title struct {
				MinLength int `yaml:"min_length" env-required:"true"`
				MaxLength int `yaml:"max_length" env-required:"true"`
			} `yaml:"title" env-required:"true"`
			Descriptions struct {
				MinLength int `yaml:"min_length" env-required:"true"`
				MaxLength int `yaml:"max_length" env-required:"true"`
			} `yaml:"descriptions" env-required:"true"`
			DateReleased struct {
				MinValue struct {
					Year  int `yaml:"year"  env-required:"true"`
					Month int `yaml:"month"  env-required:"true"`
					Day   int `yaml:"day"  env-required:"true"`
				} `yaml:"min_value" env-required:"true"`
			} `yaml:"date_released" env-required:"true"`
			Duraion struct {
				MinLength int `yaml:"min_length" env-required:"true"`
				MaxLength int `yaml:"max_length" env-required:"true"`
			} `yaml:"duration" env-required:"true"`
			EpisodeNumber struct {
				MaxValue int `yaml:"max_value" env-required:"true"`
			} `yaml:"episode_number" env-required:"true"`
			SeasonNumber struct {
				MaxValue int `yaml:"max_value" env-required:"true"`
			} `yaml:"season_number" env-required:"true"`
		} `yaml:"film" env-required:"true"`

		Series struct {
			Title struct {
				MinLength int `yaml:"min_length" env-required:"true"`
				MaxLength int `yaml:"max_length" env-required:"true"`
			} `yaml:"title" env-required:"true"`
			Descriptions struct {
				MinLength int `yaml:"min_length" env-required:"true"`
				MaxLength int `yaml:"max_length" env-required:"true"`
			} `yaml:"descriptions" env-required:"true"`
			DateStarted struct {
				MinValue struct {
					Year  int `yaml:"year"  env-required:"true"`
					Month int `yaml:"month"  env-required:"true"`
					Day   int `yaml:"day"  env-required:"true"`
				} `yaml:"min_value" env-required:"true"`
			} `yaml:"date_started" env-required:"true"`
			DateEnded struct {
				MinValue struct {
					Year  int `yaml:"year"  env-required:"true"`
					Month int `yaml:"month"  env-required:"true"`
					Day   int `yaml:"day"  env-required:"true"`
				} `yaml:"min_value" env-required:"true"`
			} `yaml:"date_ended" env-required:"true"`
		} `yaml:"series" env-required:"true"`
	} `yaml:"validation" env-required:"true"`
}
