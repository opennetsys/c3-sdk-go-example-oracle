function getLastEvent(instance) {
  return new Promise((resolve, reject) => {
    instance.allEvents()
    .watch((error, log) => {
      if (error) return reject(error)
      resolve(log)
    })
  })
}

module.exports = getLastEvent
