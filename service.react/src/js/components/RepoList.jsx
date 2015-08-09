import React, { PropTypes } from 'react';
import Immutable from 'immutable';
import RepoListItem from './RepoListItem';

class RepoList extends React.Component {

  static propTypes = {
    repos: PropTypes.instanceOf(Immutable.List).isRequired
  }

  render() {
    const {
      props: { repos, username }
    } = this;

    return (
      <ol>
        {repos.map(repo => <RepoListItem key={repo.get('name')} {...{ repo, username }} />)}
      </ol>
    );
  }
}

export default RepoList;
