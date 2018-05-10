import axios from 'axios'

export function getData (params) {
  return axios({
    url: '/api/hotspots',
    method: 'get',
    params
  })
}
