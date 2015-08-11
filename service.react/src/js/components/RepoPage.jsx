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
          <p>loading</p> :
          <p>
            <h3>
              {repo.get('name')}
              <small>
                &nbsp;({repo.get('stargazers_count')}&nbsp;
                  <span className="glyphicon glyphicon-star-empty"></span>)
              </small>
            </h3>
            <span>by </span>
            <Link to={`/${username}`}>{username}</Link>
            <br />
            <br />
            {repo.get('description')}
          </p>
        }

      </div>
    );
  }
}
