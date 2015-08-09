import React from 'react';
import { connect } from 'react-redux';
import { prepareRoute } from '../decorators';
import * as UserActionCreators from '../actions/user';
import { Link } from 'react-router';

@prepareRoute(async function ({ store }) {
  return await store.dispatch(UserActionCreators.getUsers());
})
@connect(({ User }) => ({ User }))
class UserPage extends React.Component {

  render () {
    const { User } = this.props;
    const users = User.get('all');

    return (
      <div>
        <h3>Github</h3>
          {users && users.map(user =>
            <Link to={`/${user.get('name')}`}>
              {user.get('name')}
            </Link>
          )}
      </div>
    );
  }
}

export default UserPage;
