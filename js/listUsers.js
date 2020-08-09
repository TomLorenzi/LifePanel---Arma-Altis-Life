'use strict';

const e = React.createElement;

class ListUsers extends React.Component {
  render() {
    return (
        <li id="info" className="timeline-item bg-white rounded ml-3 p-4 shadow">
        <div className="timeline-arrow" />
        <div className="col-lg-12">
          <a data-toggle="collapse" href={'#Cas' + this.props.userId} role="button" aria-expanded="false" aria-controls={'Cas' + this.props.userId} className="btn btn-primary btn-block py-2 shadow-sm with-chevron">
            <p className="d-flex align-items-center justify-content-between mb-0 px-3 py-2"><strong className="text-uppercase">{this.props.uLogin}</strong><i className="fa fa-angle-down" /></p>
          </a>
          <div id={'Cas' + this.props.userId} className="collapse shadow-sm">
            <div className="card">
              <div className="card-body">
                <h2 className="h5 mb-2">Informations</h2>
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
                <p className="lead mt-4">
                    <a className="btn btn-primary btn-sm" href={"./modifyUser.html?idUser=" + this.props.userId} role="button">Modify user</a>
                    <a className="btn btn-primary btn-sm btn-suppr-evt-h ml-3" data-hid={this.props.userId} role="button" style={{color :'white'}} onClick={deleteUser} data-hid={this.props.userId}>Delete user</a>
                </p>
              </div>
            </div>
          </div>
        </div>
      </li>
    );
  }
}