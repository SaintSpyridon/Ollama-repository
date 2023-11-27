import os
import re
import subprocess
import json
import speech_recognition as sr


def _generate_response(prompt):
    """
    Function to generate and parse a response from a given LLM to be parsed by Ollama.

    Params:
        prompt: str
    """
    command = f'''curl http://localhost:11434/api/generate -d '{{ \"model\": \"mistral\", \"prompt\": \"{prompt}\" }}' | jq -r '.response' | tr -d '\\n' | sed 's/\\. /\\.\\n/g' | say'''
    try:
        subprocess.run(command, shell=True, capture_output=True, text=True, check=True)
    except subprocess.CalledProcessError as e:
        pass

def _capture_audio():
    """
    Function to capture audio from the MAC.
    """
    recognizer = sr.Recognizer()
    with sr.Microphone() as source:
        # os.system('clear')
        print('say something...')
        audio = recognizer.listen(source, timeout=5)
    try:
        prompt = recognizer.recognize_google(audio)
        return prompt
    except sr.UnknownValueError:
        return ""
    except sr.RequestError as e:
        return ''


def parse():
    """
    Function to parse input and output of the LLM.
    """
    try:
        prompt = _capture_audio()
        if prompt == 'quit':
            # os.system('clear')
            os._exit(0)
        if not prompt:
            return
        _generate_response(prompt)
    except:
        pass


def main():
    while True:
        parse()


if __name__ == '__main__':
        main()
