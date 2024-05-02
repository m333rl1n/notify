package discord

import (
	"go.uber.org/multierr"

	"github.com/projectdiscovery/gologger"
	"github.com/projectdiscovery/notify/pkg/utils"
	sliceutil "github.com/projectdiscovery/utils/slice"
)

type Provider struct {
	Discord []*Options `yaml:"discord,omitempty"`
	counter int
}

type Options struct {
	ID                      string `yaml:"id,omitempty"`
	DiscordWebHookURL       string `yaml:"discord_webhook_url,omitempty"`
	DiscordWebHookUsername  string `yaml:"discord_username,omitempty"`
	DiscordWebHookAvatarURL string `yaml:"discord_avatar,omitempty"`
	DiscordThreads          bool   `yaml:"discord_threads,omitempty"`
	DiscordThreadID         string `yaml:"discord_thread_id,omitempty"`
	DiscordFormat           string `yaml:"discord_format,omitempty"`
}

func New(options []*Options, ids []string) (*Provider, error) {
	provider := &Provider{}

	for _, o := range options {
		if len(ids) == 0 || sliceutil.Contains(ids, o.ID) {
			provider.Discord = append(provider.Discord, o)
		}
	}

	provider.counter = 0

	return provider, nil
}
func (p *Provider) Send(message, CliFormat string) error {
	var errs []error
	for _, pr := range p.Discord {
		msg := utils.FormatMessage(message, utils.SelectFormat(CliFormat, pr.DiscordFormat), p.counter)
		gologger.Verbose().Msgf("discord notification sent for id: %s", msg)

		pr.SendThreaded(msg)

		gologger.Verbose().Msgf("discord notification sent for id: %s", pr.ID)
	}
	return multierr.Combine(errs...)
}
