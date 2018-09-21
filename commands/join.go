package commands

import (
	"fmt"

	"github.com/bottleneckco/radio-clerk/util"
	"github.com/bwmarrin/discordgo"
)

func join(s *discordgo.Session, m *discordgo.MessageCreate) {
	voiceState, err := util.FindUserVoiceState(s, m.Author.ID)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s you are not in a voice channel", m.Author.Mention()))
		return
	}
	channel, err := s.Channel(voiceState.ChannelID)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s error occurred: %s", m.Author.Mention(), err))
		return
	}
	voiceChannel, err := s.ChannelVoiceJoin(voiceState.GuildID, voiceState.ChannelID, false, true)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s error occurred: %s", m.Author.Mention(), err))
		return
	}
	VoiceConnection = voiceChannel
	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s joined '%s'", m.Author.Mention(), channel.Name))
}