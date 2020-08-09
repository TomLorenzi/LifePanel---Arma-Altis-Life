'use strict';

const e = React.createElement;

class Jumbotron extends React.Component {
  render() {
    return (
      <div className="px-3 py-3">
        <div className="jumbotron" style={{textAlign:'justify'}}>
            <h1 className="display-4">Le festival en 2020</h1>
            <p className="lead py-2">A l'aube de sa 73e édition, le Festival international du film de Cannes demeure l'un des événements les plus médiatisés au monde et le plus important festival de cinéma du point de vue du rayonnement international. Pour durer, le Festival a dû rester fidèle à sa vocation fondatrice qui était de révéler et mettre en valeur des œuvres de qualité pour servir l'évolution du cinéma.</p>
            <p className="lead" style={{fontStyle:'italic'}}> - Le travail le plus important mais aussi le moins visible, est réalisé par les "têtes chercheuses" de l'équipe qui parcourent le monde et les festivals chaque année pour dénicher les réalisateurs prometteurs.</p>
            <p className="lead" style={{fontStyle:'italic'}}> - La "sélection officielle" c’est toute la diversité de la création cinématographique qui est mise en valeur à travers différentes sélections qui ont chacune leur identité.</p>
            <hr className="my-4"></hr>
            <p className="lead">
              <a className="btn btn-primary" href="./presentation.html" role="button">En savoir plus</a>
            </p>
        </div>
      </div>
    );
  }
}

ReactDOM.render(e(Jumbotron), document.querySelector('#jumbotron'));

// ReactDOM.render(e(Menu), document.querySelector('#menu_bis'));
