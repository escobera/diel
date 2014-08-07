(function() {

  var App = Ember.Application.create();

  App.Estudante = DS.Model.extend({
    nome: DS.attr('string'),
    matricula: DS.attr('string'),
    created: DS.attr('date'),
  });

  App.Router.map(function() {
    this.resource('estudantes', { path: '/' })
  });

  App.EstudantesRoute = Ember.Route.extend({
    model: function() {
      return this.store.find('estudante');
    }
  })

})();