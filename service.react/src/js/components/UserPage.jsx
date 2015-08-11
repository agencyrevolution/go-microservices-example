import React from 'react';
import { connect } from 'react-redux';
import { prepareRoute } from '../decorators';
import * as UserActionCreators from '../actions/user';
import { Link } from 'react-router';

@prepareRoute(async function ({ store, params: { username } }) {
  return await * [
    store.dispatch(UserActionCreators.getOneByUsername(username))
  ];
})
@connect(({ User }) => ({ User }))
class UserPage extends React.Component {

  render () {
    const {
      props: {
        User,
        params: { username }
      }
    } = this;

    const user = User.get(username);

    return (
      <div>
        <h3>{user ? username : 'Loading...'}</h3>
        <div>
          {user ? user.get('description') : 'Loading...'}
        </div>
        <Link to={`/${username}/repos`}>
          View repositories
        </Link>
      </div>
    );
  }
}

export default UserPage;
