export default class Request {
  static baseURL = '/api'

  /* AWAITABLES */
  static getById = (controller, param) => Request._send({ method: 'GET', controller, param })
  static getAll = (controller) => Request._send({ method: 'GET', controller })
  static postNew = (controller, data) => Request._send({ method: 'POST', controller, data })
  static putById = (controller, param, data) => Request._send({ method: 'PUT', controller, param, data })
  static deleteById = (controller, param) => Request._send({ method: 'DELETE', controller, param })

  /* RESULT IN JSON ON RESOLVE */
  static _send({ method, controller, param, data }) {
    if (!method || !['GET', 'POST', 'PUT', 'DELETE'].includes(method) || !controller) return;

    let url = param ? `${Request.baseURL}/${controller}/${param}` : `${Request.baseURL}/${controller}`
    let options = {
      method,
      cache: 'no-cache',
      credentials: 'same-origin',
      headers: { 'Content-Type': 'application/json' },
    }

    if (method === 'POST' && data) options = { ...options, body: JSON.stringify(data) }

    return new Promise((resolve, reject) =>
      fetch(url, options)
        .then(res => res.json())
        .then(resolve)
        .catch(err => reject(err.message)))
  }
}