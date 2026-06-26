
const GAME_NAME_MIN_LENGTH = 3
const GAME_NAME_MAX_LENGTH = 16
const TAG_LINE_MIN_LENGTH = 3
const TAG_LINE_MAX_LENGTH = 5
const DELIMITER = "#"

/**
 * Takes a string and interprets it as a Riot account game name and tag line, 
 * where the tag line is separated by the game name by "#". If there is an
 * error with the input string, then the error value returned is a string
 * describing the issue. If the `gameName` and `tagLine` were properly parsed
 * then `error` will be `null`.
 * 
 * Note that the returned values will not necessarily be URI-encoded.
 */
export function validateRiotAccount(input: string): [
    gameName: string, 
    tagLine: string, 
    error: string | null
] {
    let gameName = ""
    let tagLine = ""
    let foundDelimiter = false

    for (let i = 0; i < input.length; i++) {
        if (foundDelimiter) {
            tagLine += input[i]
            continue
        }
        if (input[i] === DELIMITER) {
            foundDelimiter = true
            continue
        }
        gameName += input[i]
    }

    if (gameName.length > GAME_NAME_MAX_LENGTH) {
        return [ 
            "", "", 
            `The game name cannot have more than ${GAME_NAME_MAX_LENGTH} characters`,
        ]
    }
    if (gameName.length < GAME_NAME_MIN_LENGTH) {
        return [
            "", "",
            `The game name cannot have fewer than ${GAME_NAME_MIN_LENGTH} characters`,
        ]
    }
    if (!foundDelimiter) {
        return [ "", "", "Make sure to include a tag line like \"#NA1\"" ]
    }
    if (tagLine.length > TAG_LINE_MAX_LENGTH) {
        return [ 
            "", "", 
            `The tag line cannot have more than ${TAG_LINE_MAX_LENGTH} characters`,
        ]
    }
    if (tagLine.length < TAG_LINE_MIN_LENGTH) {
        return [
            "", "",
            `The tag line cannot have fewer than ${TAG_LINE_MIN_LENGTH} characters`,
        ]
    }

    return [ gameName, tagLine, null ]
}
