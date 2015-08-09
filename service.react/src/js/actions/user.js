import ActionTypes from '../consts/ActionTypes';

export function getOneByUsername(username) {
  return async api => ({
    type: ActionTypes.User.getOneByUsername,
    username,
    res: await api(`/users/${username}`)
  });
}

export function getUsers() {
  return async api => ({
    type: ActionTypes.User.getUsers,
    res: await api(`/users`)
  });
}
