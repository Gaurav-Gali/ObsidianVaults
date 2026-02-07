[Source](https://learn.deeplearning.ai/courses/building-ai-voice-agents-for-production/lesson/fizgb/introduction)

### What is voice agents?
Voice agents combine speech and reasoning abilities of foundation models to deliver real-time, human like voice interactions

**Approaches for building voice agents** :
1. Speech to Speech approach
	1. Simpler to implement
	2. Less flexibility and control over the agent's behaviour
2. Pipeline approach
	1. Speech to Text (input)
	2. LLM or Agentic workflow (process)
	3. Text to Speech (output)

**Components of the pipeline approach**
1. ==ASR== (Automatic Speech Recognition) a.k.a ==STT== (Speech To Text)
	1. This transcribes the given audio into text
	2. ==VAD== (Voice Activity Detection)
		1. This detects the presence and absence of human speech in audio
	3. ==EOU== (End Of Utterance)
		1. This detects weather the user has finished their turn of speaking
2. ==LLM== and ==LLM Agent==
	1. This generates the response to the transcribed audio
3. ==TTS== (Text To Speech) a.k.a Speech Synthesis
	1. This generates natural and intelligible speech from text