package prompts

const GET_IMPORTANT_BASE_PROMPT_INFO = `
Please extract any important information from the user's prompt below:

<USER_PROMPT>
%s
</USER_PROMPT>

Just write down any information that wouldn't be covered in an outline.
Please use the below template for formatting your response.
This would be things like instructions for chapter length, overall vision, instructions for formatting, etc.
(Don't use the xml tags though - those are for example only)

<EXAMPLE>
# Important Additional Context
- Important point 1
- Important point 2
</EXAMPLE>

Do NOT write the outline itself, just some extra context. Keep your responses short.

`

const STORY_ELEMENTS_PROMPT = `I'm working on writing a fictional story, and I'd like your help writing out the story elements.
Here's the prompt for my story.
<PROMPT>
%s
</PROMPT>
Please make your response have the following format:
<RESPONSE_TEMPLATE>
# Story Title
## Genre
- **Category**: (e.g., romance, mystery, science fiction, fantasy, horror)
## Theme
- **Central Idea or Message**:
## Pacing
- **Speed**: (e.g., slow, fast)
## Style
- **Language Use**: (e.g., sentence structure, vocabulary, tone, figurative language)
## Plot
- **Exposition**:
- **Rising Action**:
- **Climax**:
- **Falling Action**:
- **Resolution**:
## Setting
### Setting 1
- **Time**: (e.g., present day, future, past)
- **Location**: (e.g., city, countryside, another planet)
- **Culture**: (e.g., modern, medieval, alien)
- **Mood**: (e.g., gloomy, high-tech, dystopian)
(Repeat the above structure for additional settings)
## Conflict
- **Type**: (e.g., internal, external)
- **Description**:
## Symbolism
### Symbol 1
- **Symbol**:
- **Meaning**:
(Repeat the above structure for additional symbols)
## Characters
### Main Character(s)
#### Main Character 1
- **Name**:
- **Physical Description**:
- **Personality**:
- **Background**:
- **Motivation**:
(Repeat the above structure for additional main characters)
### Supporting Characters
#### Character 1
- **Name**:
- **Physical Description**:
- **Personality**:
- **Background**:
- **Role in the story**:
#### Character 2
- **Name**:
- **Physical Description**:
- **Personality**:
- **Background**:
- **Role in the story**:
#### Character 3
- **Name**:
- **Physical Description**:
- **Personality**:
- **Background**:
- **Role in the story**:
#### Character 4
- **Name**:
- **Physical Description**:
- **Personality**:
- **Background**:
- **Role in the story**:
#### Character 5
- **Name**:
- **Physical Description**:
- **Personality**:
- **Background**:
- **Role in the story**:
#### Character 6
- **Name**:
- **Physical Description**:
- **Personality**:
- **Background**:
- **Role in the story**:
#### Character 7
- **Name**:
- **Physical Description**:
- **Personality**:
- **Background**:
- **Role in the story**:
#### Character 8
- **Name**:
- **Physical Description**:
- **Personality**:
- **Background**:
- **Role in the story**:
(Repeat the above structure for additional supporting character)
</RESPONSE_TEMPLATE>
Of course, don't include the XML tags - those are just to indicate the example.
Also, the items in parenthesis are just to give you a better idea of what to write about, and should also be omitted from your response.`

const INITIAL_OUTLINE_PROMPT = `
Please write a markdown formatted outline based on the following prompt:

<PROMPT>
%s
</PROMPT>

<ELEMENTS>
%s
</ELEMENTS>

As you write, remember to ask yourself the following questions:
    - What is the conflict?
    - Who are the characters (at least two characters)?
    - What do the characters mean to each other?
    - Where are we located?
    - What are the stakes (is it high, is it low, what is at stake here)?
    - What is the goal or solution to the conflict?

Don't answer these questions directly, instead make your outline implicitly answer them. (Show, don't tell)

Please keep your outline clear as to what content is in what chapter.
Make sure to add lots of detail as you write.

Also, include information about the different characters, and how they change over the course of the story.
We want to have rich and complex character development!

Also limit the chapter count to %d`

const CHAPTER_COUNT_PROMPT = `
<OUTLINE>
%s
</OUTLINE>
Provide the total number of chapters from the outline above in this exact format: {"totalChapters": <number>}
No additional text or explanations. No code blocks. Strictly the format and only the JSON response.`

const CHAPTER_GENERATION_INTRO = `
You are a great fiction writer, and you're working on a great new story. 
You're working on a new novel, and you want to produce a quality output.
Here is your outline:
<OUTLINE>
%s
</OUTLINE>
`
const CHAPTER_OUTLINE_PROMPT = `
Please generate an outline for chapter %d based on the provided outline.

<OUTLINE>
%s
</OUTLINE>

As you write, keep the following in mind:
    - What is the conflict?
    - Who are the characters (at least two characters)?
    - What do the characters mean to each other?
    - Where are we located?
    - What are the stakes (is it high, is it low, what is at stake here)?
    - What is the goal or solution to the conflict?

Remember to follow the provided outline when creating your chapter outline.

Don't answer these questions directly, instead make your outline implicitly answer them. (Show, don't tell)

Please break your response into scenes, which each have the following format (please repeat the scene format for each scene in the chapter (min of 3):

# Chapter %d

## Scene: [Brief Scene Title]

- **Characters & Setting:**
  - Character: [Character Name] - [Brief Description]
  - Location: [Scene Location]
  - Time: [When the scene takes place]

- **Conflict & Tone:**
  - Conflict: [Type & Description]
  - Tone: [Emotional tone]

- **Key Events & Dialogue:**
  - [Briefly describe important events, actions, or dialogue]

- **Literary Devices:**
  - [Foreshadowing, symbolism, or other devices, if any]

- **Resolution & Lead-in:**
  - [How the scene ends and connects to the next one]

Again, don't write the chapter itself, just create a detailed outline of the chapter.  

Make sure your chapter has a markdown-formatted name!
`

const CHAPTER_SUMMARY_INTRO = "You are a helpful AI Assistant. Answer the user's prompts to the best of your abilities."

const CHAPTER_SUMMARY_PROMPT = `
I'm writing the next chapter in my novel (chapter %d), and I have the following written so far.

My outline:
<OUTLINE>
%s
</OUTLINE>

And what I've written in the last chapter:
<PREVIOUS_CHAPTER>
%s
</PREVIOUS_CHAPTER>

Please create a list of important summary points from the last chapter so that I know what to keep in mind as I write this chapter.
Also make sure to add a summary of the previous chapter, and focus on noting down any important plot points, and the state of the story as the chapter ends.
That way, when I'm writing, I'll know where to pick up again.

Here's some formatting guidelines:

Previous Chapter:
    - Plot:
        - Your bullet point summary here with as much detail as needed
    - Setting:
        - some stuff here
    - Characters:
        - character 1
            - info about them, from that chapter
            - if they changed, how so

Things to keep in Mind:
    - something that the previous chapter did to advance the plot, so we incorporate it into the next chapter
    - something else that is important to remember when writing the next chapter
    - another thing
    - etc.

Thank you for helping me write my story! Please only include your summary and things to keep in mind, don't write anything else.
`
const CHAPTER_GENERATION_STAGE1 = `
%s

%s

Please write the plot for chapter %d of %d based on the following chapter outline and any previous chapters.
Pay attention to the previous chapters, and make sure you both continue seamlessly from them, It's imperative that your writing connects well with the previous chapter, and flows into the next (so try to follow the outline)!

Here is my outline for this chapter:
<CHAPTER_OUTLINE>
%s
</CHAPTER_OUTLINE>

%s

As you write your work, please use the following suggestions to help you write chapter %d (make sure you only write this one):
    - Pacing: 
    - Are you skipping days at a time? Summarizing events? Don't do that, add scenes to detail them.
    - Is the story rushing over certain plot points and excessively focusing on others?
    - Flow: Does each chapter flow into the next? Does the plot make logical sense to the reader? Does it have a specific narrative structure at play? Is the narrative structure consistent throughout the story?
    - Genre: What is the genre? What language is appropriate for that genre? Do the scenes support the genre?

`

const CHAPTER_GENERATION_STAGE2 = `
%s

%s

Please write character development for the following chapter %d of %d based on the following criteria and any previous chapters.
Pay attention to the previous chapters, and make sure you both continue seamlessly from them, It's imperative that your writing connects well with the previous chapter, and flows into the next (so try to follow the outline)!

Don't take away content, instead expand upon it to make a longer and more detailed output.

For your reference, here is my outline for this chapter:
<CHAPTER_OUTLINE>
%s
</CHAPTER_OUTLINE>

%s

And here is what I have for the current chapter's plot:
<CHAPTER_PLOT>
%s
</CHAPTER_PLOT>

As a reminder to keep the following criteria in mind as you expand upon the above work:
    - Characters: Who are the characters in this chapter? What do they mean to each other? What is the situation between them? Is it a conflict? Is there tension? Is there a reason that the characters have been brought together?
    - Development: What are the goals of each character, and do they meet those goals? Do the characters change and exhibit growth? Do the goals of each character change over the story?
    - Details: How are things described? Is it repetitive? Is the word choice appropriate for the scene? Are we describing things too much or too little?

Don't answer these questions directly, instead make your writing implicitly answer them. (Show, don't tell)

Make sure that your chapter flows into the next and from the previous (if applicable).

Remember, have fun, be creative, and improve the character development of chapter %d (make sure you only write this one)!

`

const CHAPTER_GENERATION_STAGE3 = `
%s

%s

Please add dialogue the following chapter %d of %d based on the following criteria and any previous chapters.
Pay attention to the previous chapters, and make sure you both continue seamlessly from them, It's imperative that your writing connects well with the previous chapter, and flows into the next (so try to follow the outline)!

Don't take away content, instead expand upon it to make a longer and more detailed output.


%s

Here's what I have so far for this chapter:
<CHAPTER_CONTENT>
%s
</CHAPTER_CONTENT

As a reminder to keep the following criteria in mind:
    - Dialogue: Does the dialogue make sense? Is it appropriate given the situation? Does the pacing make sense for the scene E.g: (Is it fast-paced because they're running, or slow-paced because they're having a romantic dinner)? 
    - Disruptions: If the flow of dialogue is disrupted, what is the reason for that disruption? Is it a sense of urgency? What is causing the disruption? How does it affect the dialogue moving forwards? 
     - Pacing: 
       - Are you skipping days at a time? Summarizing events? Don't do that, add scenes to detail them.
       - Is the story rushing over certain plot points and excessively focusing on others?
    
Don't answer these questions directly, instead make your writing implicitly answer them. (Show, don't tell)

Make sure that your chapter flows into the next and from the previous (if applicable).

Also, please remove any headings from the outline that may still be present in the chapter.

Remember, have fun, be creative, and add dialogue to chapter %d (make sure you only write this one)!

`
const CHAPTER_REVISION_PROMPT = `
Please revise the following chapter:

<CHAPTER_CONTENT>
%s
</CHAPTER_CONTENT>

Based on the following feedback:
<FEEDBACK>
%s
</FEEDBACK>
Do not reflect on the revisions, just write the improved chapter that addresses the feedback and prompt criteria.  
Remember not to include any author notes.`

const CRITIC_OUTLINE_INTRO = "You are a helpful AI Assistant. Answer the user's prompts to the best of your abilities."

const CRITIC_OUTLINE_PROMPT = `
Please critique the following outline - make sure to provide constructive criticism on how it can be improved and point out any problems with it.

<OUTLINE>
%s
</OUTLINE>

As you revise, consider the following criteria:
    - Pacing: Is the story rushing over certain plot points and excessively focusing on others?
    - Details: How are things described? Is it repetitive? Is the word choice appropriate for the scene? Are we describing things too much or too little?
    - Flow: Does each chapter flow into the next? Does the plot make logical sense to the reader? Does it have a specific narrative structure at play? Is the narrative structure consistent throughout the story?
    - Genre: What is the genre? What language is appropriate for that genre? Do the scenes support the genre?

Also, please check if the outline is written chapter-by-chapter, not in sections spanning multiple chapters or subsections.
It should be very clear which chapter is which, and the content in each chapter.`

const CRITIC_CHAPTER_INTRO = "You are a helpful AI Assistant. Answer the user's prompts to the best of your abilities."

const CRITIC_CHAPTER_PROMPT = `
<CHAPTER>
%s
</CHAPTER>

Please give feedback on the above chapter based on the following criteria:
    - Pacing: Is the story rushing over certain plot points and excessively focusing on others?
    - Details: How are things described? Is it repetitive? Is the word choice appropriate for the scene? Are we describing things too much or too little?
    - Flow: Does each chapter flow into the next? Does the plot make logical sense to the reader? Does it have a specific narrative structure at play? Is the narrative structure consistent throughout the story?
    - Genre: What is the genre? What language is appropriate for that genre? Do the scenes support the genre?
    
    - Characters: Who are the characters in this chapter? What do they mean to each other? What is the situation between them? Is it a conflict? Is there tension? Is there a reason that the characters have been brought together?
    - Development:  What are the goals of each character, and do they meet those goals? Do the characters change and exhibit growth? Do the goals of each character change over the story?
    
    - Dialogue: Does the dialogue make sense? Is it appropriate given the situation? Does the pacing make sense for the scene E.g: (Is it fast-paced because they're running, or slow-paced because they're having a romantic dinner)? 
    - Disruptions: If the flow of dialogue is disrupted, what is the reason for that disruption? Is it a sense of urgency? What is causing the disruption? How does it affect the dialogue moving forwards? 
`
const OUTLINE_REVISION_PROMPT = `
Please revise the following outline:
<OUTLINE>
%s
</OUTLINE>

Based on the following feedback:
<FEEDBACK>
%s
</FEEDBACK>

Remember to expand upon your outline and add content to make it as best as it can be!


As you write, keep the following in mind:
    - What is the conflict?
    - Who are the characters (at least two characters)?
    - What do the characters mean to each other?
    - Where are we located?
    - What are the stakes (is it high, is it low, what is at stake here)?
    - What is the goal or solution to the conflict?


Please keep your outline clear as to what content is in what chapter.
Make sure to add lots of detail as you write.

Don't answer these questions directly, instead make your writing implicitly answer them. (Show, don't tell)
`

const SCRUB_PROMPT = `
<CHAPTER>
%s
</CHAPTER>

Given the above chapter, clean it up for publication by:
- Removing outlines, chapter/scene markers, --- dividers, and editorial comments
- Removing all markdown formatting (including ***, **, *, and other emphasis markers)
- Keeping only the finished story prose
- Not using XML tags in your output (they're only for this instruction)

Output only the clean text with no commentary, as this will be the final print version.
`
const PROMPT_GENERATION_INTRO = `You are an expert prompt engineer specializing in creative storytelling. Your task is to generate prompts that will be used by an LLM to create engaging stories.`

const PROMPT_GENERATION_PROMPT = `Generate %d prompts for an LLM to generate a story based on this idea: %s

Your task is to create prompts that will be used to generate stories.

Format each prompt as a numbered list (1., 2., 3., etc.).

Example prompts to follow:
%s

Do NOT reuse these previously used ideas:
%s

Remember the target audience characteristics for these stories are:
%s`
