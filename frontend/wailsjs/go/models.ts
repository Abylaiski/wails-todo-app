export namespace main {
	
	export class Task {
	    id: number;
	    title: string;
	    completed: boolean;
	    priority: string;
	    due_date: string;
	
	    static createFrom(source: any = {}) {
	        return new Task(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.title = source["title"];
	        this.completed = source["completed"];
	        this.priority = source["priority"];
	        this.due_date = source["due_date"];
	    }
	}

}

