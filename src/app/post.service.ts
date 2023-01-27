import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

const endpoint = 'https://jsonplaceholder.typicode.com/posts';
@Injectable({
  providedIn: 'root',
})
export class PostService {
  constructor(private httpService: HttpClient) {}
  getAllPosts(): Observable<any> {
    return this.httpService.get(endpoint);
  }
  getData(){
    return this.httpService.get('https://jsonplaceholder.typicode.com/todos');
  }
}