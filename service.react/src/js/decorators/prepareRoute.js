'use strict';

import React, { PropTypes } from 'react';

export default function prepareRoute(prepareFn) {

  return DecoratedComponent =>
    class PrepareRouteDecorator extends React.Component {

      static prepareRoute = prepareFn

      static contextTypes = {
        store: PropTypes.object.isRequired
      }

      render() {
        return (
          <DecoratedComponent {...this.props} />
        );
      }

      componentDidMount() {
        const {
          context: { store },
          props: { params, location }
        } = this;

        if (__FIRST__) {
          __FIRST__ = false;
          return;
        }

        prepareFn({ store, params, location });
      }
    };
}
