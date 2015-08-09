import React from 'react';
import { connect } from 'react-redux';
import { prepareRoute } from '../decorators';
import * as RepoActionCreators from '../actions/repo';
import { Link } from 'react-router';

@prepareRoute(async function ({ store, params: { username, reponame } }) {
  return await store.dispatch(RepoActionCreators.getOne({ username, reponame }));
})
@connect(({ Repo }) => ({ Repo }))
export default class RepoPage extends React.Component {

  render() {
    const {
      props: {
        Repo,
        params: { username, reponame }
      }
    } = this;

    const repo = Repo.get(`users/${username}/${reponame}`);

    return (
      <div>
        { !repo ?
          'loading' :
          <p>
            <big><strong>{repo.get('name')}</strong></big>:&nbsp;
            <small>
              ({repo.get('stargazers_count')}&nbsp;
                <span className="glyphicon glyphicon-star-empty"></span>)
            </small> by&nbsp;
            <Link to={`/${username}`}>{username}</Link>
            <br />
            {repo.get('description')}
          </p>
        }

      </div>
    );
  }
}
