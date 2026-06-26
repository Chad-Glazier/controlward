

//TODO: Implement.
//
// Makes a call to the server to get the PUUID associated with the given game
// name and tag line. If no matching player account was found, then null is 
// returned.
async function getPuuid(
    gameName: string, tagLine: string
): Promise<string | null> {
    await new Promise(resolve => setTimeout(resolve, 1000));
    if (Math.random() > 0.5) {
        return "SAMPLE_PUUID"
    } else {
        return null
    }
}

export default getPuuid
