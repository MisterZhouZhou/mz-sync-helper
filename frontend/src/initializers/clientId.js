import { v4 as UUIDV4 } from 'uuid';

let clientId = localStorage.getItem('clientId')

if (clientId?.length !== 36) {
  clientId = UUIDV4();
  localStorage.setItem('clientId', clientId)
}

export { clientId }
