const { By, Builder, until, Browser } = require('selenium-webdriver');
const chrome = require('selenium-webdriver/chrome');
const assert = require("assert");

const FindElements = async (driver) => {
	return {
		num1: await driver.findElement(By.name("num1")),
		num2: await driver.findElement(By.name("num2")),
		select: await driver.findElement(By.name("operacion")),
		button: await driver.findElement(By.css("button[type='submit']"))
	}
}

describe('First script', function () {
	this.timeout(0);

	let driver;

	before(async function () {
		this.timeout(60000);
		let options = new chrome.Options();
		options.addArguments('--no-sandbox');
		options.addArguments('--disable-dev-shm-usage');
		options.addArguments('--disable-gpu');
		options.addArguments('--disable-software-rasterizer');
		options.addArguments('--disable-extensions');
		options.addArguments('--window-size=1920,1080');

		driver = await new Builder().forBrowser('chrome').setChromeOptions(options).build();
	});

	after(async () => await driver.quit());

	const cases = [
		["2", "3", "sumar", "5"],
		["5", "2", "restar", "3"],
		["4", "6", "multiplicar", "24"],
		["10", "2", "dividir", "5"],
		["5", "0", "dividir", "error: Division by zero"],
		["abc", "def", "sumar", "error: enter valid numbers"],
	];

	cases.forEach((val) => {
		it(`calc ${val[0]} ${val[2]} ${val[1]} -> ${val[3]}`, async function () {
			await driver.get('http://localhost:8000');

			const { num1, num2, select, button } = await FindElements(driver);

			await num1.sendKeys(val[0]);
			await num2.sendKeys(val[1]);
			await select.findElement(By.css(`option[value="${val[2]}"]`)).click();
			await button.click();

			const res = await driver.wait(until.elementLocated(By.tagName("h2")), 100000);
			const num = await res.getText();

			assert.ok(num.includes(val[3]));
		});
	})
});