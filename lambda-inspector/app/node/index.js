exports.handler =  async function(event, context) {
  console.log("Who`s bad?")
  console.log("Fake it till you make it?")
  console.log("EVENT: \n" + JSON.stringify(event, null, 2))
  return context.logStreamName
}
