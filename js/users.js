'use strict';

const e = React.createElement;

class Users extends React.Component {
  render() {
    return (
        <li id="info" className="timeline-item bg-white rounded ml-3 p-4 shadow">
        <div className="timeline-arrow" />
        <div className="col-lg-12">
          <div id={'Cas' + this.props.userId}>
            <div className="card">
              <div className="card-body">
                <h2 className="h5 mb-2">Informations for {this.props.uLogin}</h2>
                <hr />
                <table style={{textAlign:'center'}} className="table table-hover">
                    <thead>
                    <tr className="table-active">
                        <th  style={{borderLeft: '3px solid #dadada'}} scope="col"><i className="fas fa-id-card mr-2"></i>User ID</th>
                        <th  scope="col"><i className="fas fa-id-badge mr-2"> </i>Login</th>
                        <th  scope="col"><i className="fas fa-lock mr-2"></i>Permissions</th>
                        
                    </tr>
                    </thead>
                    <tbody>
                    <tr>
                        <td style={{border: '3px solid #dadada'}}>{ this.props.userId}</td>
                        <td style={{border: '3px solid #dadada'}}>{ this.props.uLogin}</td>
                        <td style={{border: '3px solid #dadada'}}>{ this.props.uperms}</td>
                    </tr>
                    </tbody>
                </table> 
              </div>
            </div>
          </div>
        </div>
      </li>
    );
  }
}