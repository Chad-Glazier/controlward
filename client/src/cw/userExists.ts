

//TODO: Implement.
//
// Should also make it return the PUUID and then use PUUID endpoints to minimize
// Riot API calls.
async function userExists(gameName: string, tagLine: string): Promise<boolean> {
    await new Promise(resolve => setTimeout(resolve, 1000));
    console.log(gameName, tagLine)
    return Math.random() > 0.5
}

export default userExists
