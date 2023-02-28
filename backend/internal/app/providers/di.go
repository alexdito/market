package providers

import (
	"application/internal/app/common"
	"application/internal/app/http/v1/controller"
	"application/internal/app/repository"
	"application/internal/app/src/service"
	"sync"
)

type repositoryDI struct {
	navigationRepository repository.NavigationRepository
	cardRepository       repository.CardRepository
}

func (repositoryDI *repositoryDI) bind(app *common.App) {
	repositoryDI.navigationRepository = repository.NewNavigationRepository(app.DataBase)
	repositoryDI.cardRepository = repository.NewCardRepository(app.DataBase)
}

type serviceDI struct {
	navigationService service.NavigationService
	cardService       service.CardService
}

func (serviceDI *serviceDI) bind(mu *sync.RWMutex, repositories *repositoryDI) {
	serviceDI.navigationService = service.NewNavigationService(mu, repositories.navigationRepository)
	serviceDI.cardService = service.NewCardService(mu, repositories.cardRepository)
}

type ControllerDI struct {
	NavigationController controller.NavigationController
	CardController       controller.CardController
}

func (controllerDI *ControllerDI) bind(services *serviceDI) {
	controllerDI.NavigationController = controller.NewNavigationController(services.navigationService)
	controllerDI.CardController = controller.NewCardController(services.cardService)
}

func BuildDI(app *common.App) *ControllerDI {
	var repoDI repositoryDI
	var serviceDI serviceDI
	var controllerDI ControllerDI

	mu := new(sync.RWMutex)

	repoDI.bind(app)
	serviceDI.bind(mu, &repoDI)
	controllerDI.bind(&serviceDI)

	return &controllerDI
}
