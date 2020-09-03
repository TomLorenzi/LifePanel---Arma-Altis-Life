'use strict';

const e = React.createElement;

class Logs extends React.Component {
  render() {
    return (
        <li id="info" className="timeline-item bg-white rounded ml-3 p-4 shadow">
        <div className="timeline-arrow" />
          <div id={'Cas' + this.props.lId}>
                <table style={{textAlign:'center'}} className="table table-hover">
                    <thead>
                    <tr className="table-active">
                        <th  style={{borderLeft: '3px solid #dadada'}} scope="col"><i className="fas fa-id-card mr-2"></i>Log id</th>
                        <th  scope="col">Player</th>
                        <th  scope="col">Action</th>
                        <th  scope="col">Admin</th>
                        <th  style={{borderRight: '3px solid #dadada'}} scope="col">Date</th>
                    </tr>
                    </thead>
                    <tbody>
                    <tr>
                        <td style={{border: '3px solid #dadada'}}>{ this.props.lId}</td>
                        <td style={{border: '3px solid #dadada'}}>{ this.props.lPlayer}</td>
                        <td style={{border: '3px solid #dadada'}}>{ this.props.lAction}</td>
                        <td style={{border: '3px solid #dadada'}}>{ this.props.lAdmin}</td>
                        <td style={{border: '3px solid #dadada'}}>{ this.props.lDate}</td>
                    </tr>
                    </tbody>
                </table> 
        </div>
      </li>
    );
  }
}