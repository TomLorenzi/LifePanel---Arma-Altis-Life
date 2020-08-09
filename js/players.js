'use strict';

const e = React.createElement;

class Players extends React.Component {
  render() {
    return (
        <li id="info" className="timeline-item bg-white rounded ml-3 p-4 shadow">
        <div className="timeline-arrow" />
        <div className="col-lg-12">
          <div id={'Cas' + this.props.playerId}>
            <div className="card">
              <div className="card-body">
                <h2 className="h5 mb-2">Informations for user with PID : {this.props.pPid}</h2>
                <hr />
                <table style={{textAlign:'center'}} className="table table-hover">
                    <thead>
                    <tr className="table-active">
                        <th  style={{borderLeft: '3px solid #dadada'}} scope="col"><i className="fas fa-id-card mr-2"></i>User ID</th>
                        <th  scope="col"><i className="fas fa-id-badge mr-2"> </i>PID</th>
                        <th  scope="col"><i className="fas fa-id-badge mr-2"> </i>Nom</th>
                        <th  scope="col"><i className="fas fa-credit-card mr-2"></i>Cash</th>
                        <th  scope="col"><i className="fas fa-university mr-2"></i>Bank</th>
                        <th  scope="col"><i className="fas fa-male mr-2"></i>Cop Level</th>
                        <th  scope="col"><i className="fas fa-user-md mr-2"></i>Medic Level</th>
                        <th  scope="col"><i className="fas fa-user-plus mr-2"></i>Donor Level</th>
                        <th  style={{borderRight: '3px solid #dadada'}} scope="col"><i className="fas fa-user-secret mr-2"></i>Admin Level</th>
                        
                    </tr>
                    </thead>
                    <tbody>
                    <tr>
                        <td style={{border: '3px solid #dadada'}}>{ this.props.playerId}</td>
                        <td style={{border: '3px solid #dadada'}}>{ this.props.pPid}</td>
                        <td style={{border: '3px solid #dadada'}}>{ this.props.pName}</td>
                        <td style={{border: '3px solid #dadada'}}>{ this.props.pcash}</td>
                        <td style={{border: '3px solid #dadada'}}>{ this.props.pbank}</td>
                        <td style={{border: '3px solid #dadada'}}>{ this.props.pcoplevel}</td>
                        <td style={{border: '3px solid #dadada'}}>{ this.props.pmediclevel}</td>
                        <td style={{border: '3px solid #dadada'}}>{ this.props.pdonorlevel}</td>
                        <td style={{border: '3px solid #dadada'}}>{ this.props.padminlevel}</td>
                    </tr>
                    </tbody>
                </table> 
                <p className="lead mt-4">
                    
                </p>
              </div>
            </div>
          </div>
        </div>
      </li>
    );
  }
}