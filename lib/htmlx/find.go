package htmlx

import (
	"golang.org/x/net/html"
	"strings"
)

// HasClass 检查节点是否具有指定class
func HasClass(class string) func(n *html.Node) bool {
	return func(n *html.Node) bool {
		for _, attr := range n.Attr {
			if attr.Key == "class" && attr.Val == class {
				return true
			}
		}
		return false
	}
}

//// HasPrefixClass 检查节点是否具有指定class名前缀
//func HasPrefixClass(class string) func(n *html.Node) bool {
//	return func(n *html.Node) bool {
//		for _, attr := range n.Attr {
//			if attr.Key == "class" && attr.Val == class {
//				return true
//			}
//		}
//		return false
//	}
//}

// FindElements 递归查找所有符合check的元素
func FindElements(n *html.Node, check func(n *html.Node) bool) []*html.Node {
	var elements []*html.Node

	if n.Type == html.ElementNode && check(n) {
		elements = append(elements, n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		elements = append(elements, FindElements(c, check)...)
	}

	return elements
}

// FindElement 递归查找一个符合check的元素
func FindElement(n *html.Node, check func(n *html.Node) bool) *html.Node {
	if n.Type == html.ElementNode && check(n) {
		return n
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if found := FindElement(c, check); found != nil {
			return found
		}
	}

	return nil
}

// GetTextContent 获取元素的文本内容
func GetTextContent(n *html.Node) string {
	var result strings.Builder
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode {
			result.WriteString(c.Data)
		}
	}
	return strings.TrimSpace(result.String())
}
