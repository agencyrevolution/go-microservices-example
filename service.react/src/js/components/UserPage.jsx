import React from 'react';
import { connect } from 'react-redux';
import { prepareRoute } from '../decorators';
import * as RepoActionCreators from '../actions/repo';
import RepoList from './RepoList';

@prepareRoute(async function ({ store, params: { username } }) {
  return await * [
    store.dispatch(RepoActionCreators.getByUsername(username))
  ];
})
@connect(({ Repo, User }) => ({ Repo, User }))
class UserPage extends React.Component {

  render () {
    const {
      props: {
        Repo,
        params: { username }
      }
    } = this;

    const repos = Repo.get(`users/${username}`);

    return (
      <div>
        <h4>Repositories:</h4>
        {repos ? <RepoList {...{ repos, username }} /> : 'Loading...'}
      </div>
    );
  }
}

export default UserPage;
