import { BrowserModule } from '@angular/platform-browser';
import { Component,NgModule } from '@angular/core';
import { FormsModule, ReactiveFormsModule} from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';
import {RouterModule} from '@angular/router'
import { HomeComponent } from './home.component';
import { AppComponent } from './app.component';

@Component({
  selector: 'my-app',
  template: `
    <router-outlet></router-outlet>
  `,
})
export class App {
  name:string;
  constructor() {
    this.name = 'Angular2'
  }
}

const routes = [
  {
    path: "",
    component: HomeComponent
  },
  {
    path: "addCustomer",
    component: AppComponent
  }
]
@NgModule({
  declarations: [
    App,
    HomeComponent,
    AppComponent
  ],
  imports: [
    BrowserModule,
    HttpClientModule,
    RouterModule.forRoot(routes),
    FormsModule, 
    ReactiveFormsModule
  ],
  providers: [],
  bootstrap: [App]
})
export class AppModule { }
