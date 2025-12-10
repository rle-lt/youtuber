package prompts

const GET_IMPORTANT_BASE_PROMPT_INFO = `
Please extract any important information from the user's prompt below:

<USER_PROMPT>
%s
</USER_PROMPT>

Just write down any information that wouldn't be covered in an outline.
Please use the below template for formatting your response.
This would be things like instructions for scene length, overall vision, instructions for formatting, target runtime, etc.
(Don't use the xml tags though - those are for example only)

<EXAMPLE>
# Important Additional Context
- Important point 1
- Important point 2
</EXAMPLE>

Do NOT write the outline itself, just some extra context. Keep your responses short.
`

const STORY_ELEMENTS_PROMPT = `I'm working on writing a dramatic story for video narration, and I'd like your help defining the story elements.

Here's the prompt for my story:
<PROMPT>
%s
</PROMPT>

Please structure your response with these elements:

# Story Title

## Genre
- **Primary Genre**: (thriller, drama, sci-fi, horror, mystery, etc.)
- **Subgenre/Tone**: (psychological, action-driven, emotional, dark comedy)

## Hook
- **Opening Question**: What question drives the first 60 seconds?
- **Central Tension**: What keeps viewers watching?

## Theme
- **Core Message**: What's the story really about?
- **Emotional Journey**: What should viewers feel by the end?

## Conflict Structure
- **External Conflict**: The visible problem/antagonist
- **Internal Conflict**: Character's personal struggle
- **Stakes**: What happens if they fail?

## Setting
- **Primary Location(s)**: Where does this take place?
- **Time Period**: When does this occur?
- **Atmosphere**: What's the mood/feeling of the world?

## Protagonist
- **Name**:
- **Core Trait**: The one thing that defines them
- **What They Want**: External goal
- **What They Need**: Internal growth
- **Fatal Flaw**: What could cause their downfall?

## Supporting Characters (3-5 key characters)
### Character 1
- **Name**:
- **Role**: (ally, antagonist, victim, mentor, etc.)
- **Relationship to Protagonist**:
- **Function in Story**: Why they matter

### Character 2
- **Name**:
- **Role**:
- **Relationship to Protagonist**:
- **Function in Story**:

### Character 3
- **Name**:
- **Role**:
- **Relationship to Protagonist**:
- **Function in Story**:

### Character 4
- **Name**:
- **Role**:
- **Relationship to Protagonist**:
- **Function in Story**:

### Character 5
- **Name**:
- **Role**:
- **Relationship to Protagonist**:
- **Function in Story**:

## Plot Arc (4-Act Structure)
- **Act 1 - Setup** (25%%): What's the normal world? What disrupts it?
- **Act 2 - Rising Action** (35%%): How does the situation escalate?
- **Act 3 - Crisis** (25%%): What's the darkest moment?
- **Act 4 - Resolution** (15%%): How does it resolve? What's changed?

## Key Dramatic Moments
- **Point of No Return**: When the protagonist commits
- **Midpoint Twist**: What changes everything?
- **All Is Lost**: The lowest point
- **Climax**: The final confrontation/revelation

Keep the structure clear and focused on dramatic tension throughout.`

const INITIAL_OUTLINE_PROMPT = `
Create a detailed scene-by-scene outline for a 60-minute dramatic video story.

<PROMPT>
%s
</PROMPT>

<ELEMENTS>
%s
</ELEMENTS>

Structure your outline with 12-15 scenes following a 4-act dramatic structure:

**ACT 1: HOOK & SETUP (15 minutes, 3-4 scenes)**
- Scene 1: Cold open - most dramatic/intriguing moment
- Scene 2-3: Establish world, character, stakes
- Scene 4: Inciting incident - the point of no return

**ACT 2: ESCALATION (20 minutes, 4-5 scenes)**
- Rising tension and complications
- Midpoint: major twist or revelation
- Stakes increase, situation worsens

**ACT 3: CRISIS (20 minutes, 4-5 scenes)**
- Everything falls apart
- Protagonist's lowest point
- Forces final decision/confrontation

**ACT 4: CLIMAX & RESOLUTION (5 minutes, 2-3 scenes)**
- Final confrontation/revelation
- Resolution of conflict
- Emotional payoff and thematic statement

For each scene, include:
- Scene number and title
- Location and time
- Characters present
- What happens (action beats)
- Emotional purpose
- How it connects to next scene
- Estimated runtime (3-5 minutes)

Keep each scene focused on ONE dramatic beat while building overall tension.
Ensure each scene ends with a hook that makes viewers need to see what happens next.

REMEMVER, max limit total scenes is %d`

const SCENE_COUNT_PROMPT = `
<OUTLINE>
%s
</OUTLINE>
Provide the total number of scenes from the outline above in this exact format: {"totalChapters": <number>}
No additional text or explanations. No code blocks. Strictly the format and only the JSON response.`

const SCENE_GENERATION_INTRO = `
You are a skilled dramatic writer creating narration for video content.
Here is your story outline:
<OUTLINE>
%s
</OUTLINE>
`

const SCENE_OUTLINE_PROMPT = `
Create a detailed outline for scene %d based on the story outline.

<FULL_OUTLINE>
%s
</FULL_OUTLINE>

Use this structure:

# Scene %d: [Compelling Title]

## Core Scene Information
- **Duration**: [Target: 3-5 minutes]
- **Location**: [Specific setting with visual details]
- **Time**: [Time of day, lighting, atmosphere]
- **Characters Present**: [List with their emotional states]

## Scene Purpose
- **Plot Function**: What story information is conveyed?
- **Emotional Beat**: What should viewers feel?
- **Character Development**: How do characters change/reveal themselves?
- **Tension Level**: [Low/Building/High/Peak]

## Scene Breakdown (Beat by Beat)
1. **Opening**: How does the scene start?
2. **Conflict/Complication**: What goes wrong or what's revealed?
3. **Escalation**: How does tension build?
4. **Climax**: The peak moment of the scene
5. **Transition**: How does it lead to the next scene?

## Visual Elements
- **Key Images**: What should the main visual(s) show?
- **Atmosphere**: Color palette, lighting, mood
- **Focus**: What draws the viewer's eye?

## Dialogue & Narration Notes
- **Key Lines**: Important dialogue or narration
- **Tone**: [Tense, emotional, mysterious, urgent, etc.]
- **Pacing**: [Fast/Medium/Slow - matches scene intensity]

## Story Questions
- **Question Raised**: What makes viewers need the next scene?
- **Information Revealed**: What do we learn?
- **Mystery/Tension**: What's still unresolved?

Ensure this scene builds on what came before and propels the story forward.`

const SCENE_SUMMARY_INTRO = "You are a helpful AI Assistant. Answer the user's prompts to the best of your abilities."

const SCENE_SUMMARY_PROMPT = `
I'm writing scene %d in my video story, and I need to know what happened in the previous scene.

Story outline:
<OUTLINE>
%s
</OUTLINE>

Previous scene content:
<PREVIOUS_SCENE>
%s
</PREVIOUS_SCENE>

Please create a summary of the previous scene focusing on:

Previous Scene Summary:
    - Plot Events:
        - Key actions and decisions that occurred
    - Setting:
        - Where it took place and atmospheric details
    - Characters:
        - Who was involved and their emotional state at scene end
        - Any character development or revelations
    - Unresolved Tension:
        - Questions or conflicts left hanging

Things to Keep in Mind for Next Scene:
    - Plot threads to continue or resolve
    - Character states to maintain consistency
    - Tone or pacing considerations
    - Visual or thematic elements to carry forward

Keep the summary focused on what I need to write the next scene seamlessly.`

const SCENE_GENERATION_STAGE1 = `
 %s

%s

Write the narration script for scene %d of %d for video voiceover.

Scene outline:
<SCENE_OUTLINE>
%s
</SCENE_OUTLINE>

%s

Requirements for narration:
- Written to be READ ALOUD naturally
- Present tense for immediacy
- Short, punchy sentences for dramatic impact
- Varies sentence length to control pacing
- Include strategic pauses [PAUSE] for emphasis
- Balance description, action, and emotion
- Conversational but powerful language
- Target length: approximately 4-5 minutes when read at natural speaking pace

Structure:
1. Opening Hook (2-3 sentences): Pull viewers into the scene
2. Setup (describe setting, establish mood)
3. Action/Dialogue (what happens, what's said)
4. Emotional Beat (the feeling/realization)
5. Transition (set up next scene or leave hook)

Write ONLY the narration text for scene %d - no stage directions, no author notes, no scene headings.`

const SCENE_GENERATION_STAGE2 = `
%s

%s

Enhance the narration for scene %d of %d by deepening character moments and emotional resonance.

Scene outline:
<SCENE_OUTLINE>
%s
</SCENE_OUTLINE>

%s

Current narration:
<CURRENT_NARRATION>
%s
</CURRENT_NARRATION>

Expand the narration by:
- Adding character internal thoughts and motivations where appropriate
- Deepening emotional moments with more vivid description
- Showing character reactions through body language and expressions
- Building tension through character perspective
- Ensuring character voices and behaviors feel authentic

Keep the voiceover natural for reading aloud. Do not remove content, only enhance and expand.

Write the enhanced narration for scene %d.`

const SCENE_GENERATION_STAGE3 = `
%s

%s

Add and refine dialogue for scene %d of %d, ensuring it sounds natural when read aloud.

%s

Current narration:
<CURRENT_NARRATION>
%s
</CURRENT_NARRATION>

Enhance the narration by:
- Adding dialogue where characters interact
- Ensuring dialogue sounds natural and character-specific
- Balancing dialogue with narrative description
- Using dialogue to reveal character and advance plot
- Matching dialogue pacing to scene intensity
- Removing any remaining outline headings or structural markers

The final output should flow as pure narration/dialogue suitable for voiceover recording.

Write the final polished narration for scene %d with all dialogue integrated.`

const CRITIC_OUTLINE_INTRO = "You are a helpful AI Assistant. Answer the user's prompts to the best of your abilities."

const CRITIC_OUTLINE_PROMPT = `
Please critique the following story outline - provide constructive criticism on how it can be improved.

<OUTLINE>
%s
</OUTLINE>

Evaluate based on:
    - Dramatic Structure: Does it follow a clear arc with rising tension?
    - Pacing: Are scenes balanced in runtime and dramatic weight?
    - Hook Strength: Does it open with immediate intrigue?
    - Character Journey: Is there meaningful character development?
    - Emotional Beats: Are there clear emotional peaks and valleys?
    - Scene Clarity: Is it clear what happens in each scene and why?
    - Visual Potential: Can each scene be represented with compelling imagery?
    - Viewer Retention: Will viewers stay engaged throughout?

Also verify the outline is structured scene-by-scene with clear boundaries, not vague sections.`

const OUTLINE_REVISION_PROMPT = `
Please revise the following outline:
<OUTLINE>
%s
</OUTLINE>

Based on the following feedback:
<FEEDBACK>
%s
</FEEDBACK>

Expand the outline and add detail to address the feedback while maintaining:
- Clear scene-by-scene structure
- Strong dramatic arc with escalating tension
- Compelling hooks and emotional beats
- Visual potential for each scene
- Target runtime of 60 minutes (12-15 scenes)

As you revise, keep in mind:
- What is the central conflict?
- Who are the characters and what do they want?
- What are the stakes?
- How does tension build scene by scene?
- What's the emotional payoff?

Make your outline detailed enough to guide scene generation while keeping dramatic focus.`

const CRITIC_SCENE_INTRO = "You are a helpful AI Assistant. Answer the user's prompts to the best of your abilities."

const CRITIC_SCENE_PROMPT = `
<SCENE>
%s
</SCENE>

Provide constructive feedback on this scene's effectiveness for video narration.

Evaluate:

**Dramatic Structure:**
- Does it start with a hook?
- Does tension build throughout?
- Is there a clear emotional beat?
- Does it end with forward momentum?

**Narration Quality:**
- Does it sound natural when read aloud?
- Is the language engaging but accessible?
- Are there clichés or weak word choices?
- Does it show rather than tell?

**Pacing:**
- Any sections that drag?
- Is sentence rhythm varied?
- Does runtime match content?
- Are pauses used effectively?

**Character:**
- Are character reactions believable?
- Do we understand motivations?
- Is there meaningful character revelation?

**Emotional Impact:**
- What should viewers feel? Do they?
- Are emotional moments earned?
- Is the tone consistent?

**Visual Support:**
- Are there clear visual moments described?
- Can you picture what's happening?
- Does it support image generation needs?

Provide specific suggestions for improvement.`

const SCENE_REVISION_PROMPT = `
Please revise the following scene:

<SCENE_CONTENT>
%s
</SCENE_CONTENT>

Based on the following feedback:
<FEEDBACK>
%s
</FEEDBACK>

Do not reflect on the revisions, just write the improved scene narration that addresses the feedback.
Output only the narration text suitable for voiceover - no author notes or commentary.`

const SCRUB_PROMPT = `
<SCENE>
%s
</SCENE>

Clean up the scene narration for final video production by:
- Removing any outline markers, scene headers, or structural notes
- Removing all markdown formatting (including ***, **, *, and emphasis markers)
- Removing editorial comments or stage directions
- Keeping only the narration text meant to be read aloud
- Preserving [PAUSE] markers for dramatic timing

Output only the clean narration text ready for voiceover recording.`

const PROMPT_GENERATION_INTRO = `You are an expert at creating compelling story premises for dramatic video content. Your prompts should generate stories that hook viewers immediately and keep them engaged.`

const PROMPT_GENERATION_PROMPT = `Generate %d compelling story prompts for dramatic video narratives.

<CORE_THEME>
%s
</CORE_THEME>

Each prompt should create a story with:

**1. Immediate Hook**
- Starts with intrigue, conflict, or impossible situation
- Makes audience ask "What happens next?"
- Can be summarized in one compelling sentence

**2. Clear Protagonist**
- Someone viewers can emotionally invest in
- Has both external goal and internal need
- Faces meaningful choices

**3. High Stakes**
- Something important at risk (life, truth, identity, relationships, justice, etc.)
- Consequences matter
- Outcome uncertain

**4. Dramatic Tension**
- Built-in conflict (person vs. person, person vs. system, person vs. self)
- Escalating complications
- Clear climax potential

**5. Emotional Payoff**
- Journey has meaning beyond plot
- Theme emerges naturally
- Satisfying resolution possible

**6. Sustainable Runtime**
- Story complexity matches %d-minute length
- Enough narrative depth to fill time without padding
- Natural story beats maintain engagement throughout

Story types to consider:
- Thrillers (mystery, revelation, pursuit)
- Moral dilemmas (impossible choices, ethical conflicts)
- Survival stories (against odds, resource scarcity)
- Conspiracy/truth-seeking (uncovering secrets, exposing lies)
- Revenge/justice (righting wrongs, confronting power)
- Sacrifice (giving up something for others)
- Redemption (making up for past, second chances)
- Rescue/protection (saving the vulnerable)
- Transformation (forced change, growth through crisis)
- Underdog challenges (unlikely heroes, David vs. Goliath)

Title structure patterns to use:
%s

Example effective prompts:
%s

DO NOT REUSE these previously used concepts:
%s

Target audience characteristics:
%s

Generate prompts that create stories viewers cannot stop watching.
Format as a numbered list (1., 2., 3., etc.).`

const PROMPT_GENERATION_TO_JSON_PROMPT = `Return the prompts as a JSON array of strings.

Your response must be valid JSON only, with no markdown formatting, no code blocks, and no comments.

Remove the list numbers from the prompts.

Format needed: {"prompts":["prompt 1", "prompt 2", "prompt 3"]}

Prompts to format:
%s`
