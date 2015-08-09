import React, { PropTypes } from 'react';
import Immutable from 'immutable';
import { Link } from 'react-router';

class RepoListItem extends React.Component {

  static propTypes = {
    repo: PropTypes.instanceOf(Immutable.Map).isRequired
  }

  render() {
    const {
      props: { repo, username }
    } = this;

    return (
      <li>
        <p>
          <Link to={`/${username}/${repo.get('name')}`}>
            <big><strong>{repo.get('name')}</strong></big>:&nbsp;
          </Link>
          <small>
            ({repo.get('stargazers_count')}&nbsp;
              <span className="glyphicon glyphicon-star-empty"></span>)
          </small>
          <br />
          {repo.get('description')}
        </p>
      </li>
    );
  }
}

export default RepoListItem;
